//user logic and database impl of interaction with
//+build go1.9

package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	_id       string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"username,omitempty"`
}

type query interface {
}

type Userrepository struct {
	collection *mongo.Collection
}

//repository query to find
func (receiver *Userrepository) FindById(id string) *bson.M {
	var filter = bson.M{}
	objID, _ := primitive.ObjectIDFromHex(id)
	filter = bson.M{"_id": objID}
	var result bson.M
	cur, _ := receiver.collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	_ = cur.All(context.Background(), &result)
	return &result
}
func (receiver *Userrepository) FindByUsername(username string) *bson.M {
	var filter = bson.M{}
	//objID, _ := primitive.ObjectIDFromHex(username)
	filter = bson.M{"username": username}
	var result bson.M
	cur, _ := receiver.collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	_ = cur.All(context.Background(), &result)
	return &result
}
//
//func AddNewUser(username, pass string) bool {
//
//	newuser := &User{
//		FirstName: "god" + strconv.FormatFloat(asin, 'f', 6, 64),
//		LastName:  "king",
//		Username:  username,
//	}
//
//	res, err := db.UserCollection.InsertOne(context.Background(), newuser)
//
//}
