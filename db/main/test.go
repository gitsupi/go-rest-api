package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
	"math/rand"
	"radical.com/go-rest-api/db"
	"radical.com/go-rest-api/test/model"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	var filter = bson.M{}
	//objID, _ := primitive.ObjectIDFromHex(username)
	username := "kingmain"
	filter = bson.M{"firstname": primitive.Regex{Pattern: "god", Options: ""}}
	var result []bson.M
	cur, _ := db.UserCollection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	_ = cur.All(context.Background(), &result)
	fmt.Printf("%v\n", result)
	fmt.Printf("%v\n", len(result))

	//var person db.Person
	//bytes, _ := json.Marshal(result[0])
	//json.Unmarshal(bytes, &person)
	x := rand.Float64() * 1000 / 2
	print(x)
	asin := math.Sinh(x)
	us := &model.User{
		FirstName:   "god" + strconv.FormatFloat(asin, 'f', 2, 32),
		LastName:    "king",
		Username:    username,
		Phonenumber: "09351844321",
		Password:    db.HashAndSalt("23"),
	}
	fmt.Printf("user %v\n", us)
	res, err := db.UserCollection.InsertOne(context.Background(), us)
	fmt.Printf("res %v\n", res)
	fmt.Printf("err %v\n", err)

}
func main() {
	Test(nil)
}
