//user logic and database impl of interaction with
//+build go1.9

package model

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"radical.com/go-rest-api/db"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	FirstName   string             `json:"firstname,omitempty"`
	LastName    string             `json:"lastname,omitempty"`
	Username    string             `json:"username,omitempty"`
	Phonenumber string             `json:"phonenumber"`
	Password    string             `json:"password,omit"`
}

type query interface {
}

//repository query to find
func (rec *User) UpdateUserInfo(info map[string]string) error {
	if val, ok := info["phonenumber"]; ok {
		fmt.Print(val)
		if checkPhoneNumberisBeforeExist(val) {
			return errors.New("phone number exist before")
		}
	}
	var filter = bson.M{"phonenumber": rec.Phonenumber}
	updateM := bson.M{"$set": info}
	upsert := true
	_, er := db.UserCollection.UpdateOne(context.Background(),
		filter,
		updateM,
		&options.UpdateOptions{Upsert: &upsert})
	return er
}

func checkPhoneNumberisBeforeExist(phonenumber string) bool {
	user, _ := GetUserByPhoneNumber(phonenumber)
	fmt.Printf("%v", user)
	if !user.Id.IsZero() {
		return true
	}
	return false
}

//repository query to find
func GetUserByPhoneNumber(phonenumber string) (*User, error) {
	var filter = bson.M{}
	filter = bson.M{"phonenumber": phonenumber}
	var user User
	err := db.UserCollection.FindOne(context.Background(), filter).Decode(&user)
	return &user, err
}

//repository query to find
func GetUserById(id string) (*User, error) {
	var filter = bson.M{}
	hex, _ := primitive.ObjectIDFromHex(id)

	filter = bson.M{"_id": hex}
	var user User
	err := db.UserCollection.FindOne(context.Background(), filter).Decode(&user)
	return &user, err
}

//repository query to find
func FindById(id string) *bson.M {
	var filter = bson.M{}
	objID, _ := primitive.ObjectIDFromHex(id)
	filter = bson.M{"_id": objID}
	var result bson.M
	cur, _ := db.UserCollection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	_ = cur.All(context.Background(), &result)
	return &result
}
func FindByUsername(username string) *bson.M {
	var filter = bson.M{}
	//objID, _ := primitive.ObjectIDFromHex(username)
	filter = bson.M{"username": username}
	var result bson.M
	cur, _ := db.UserCollection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	_ = cur.All(context.Background(), &result)
	return &result
}

func InsertNewUserByPhoneNumber(phonen string) (*mongo.InsertOneResult, error) {
	newuser := &User{
		Id:          primitive.NewObjectID(),
		Phonenumber: phonen,
	}
	return db.UserCollection.InsertOne(context.Background(), newuser)

}
