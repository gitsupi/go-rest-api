package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"radical.com/go-rest-api/db"
	"time"
)

type Acount struct {
	Id             primitive.ObjectID `bson:"_id" json:"id"`
	Enteredmoneies []Enterdmoney      `json:"enteredmoneies"`
	Type           string
}

type Enterdmoney struct {
	Amount int
	Time   time.Time
}

func (receiver *Acount) AddNewEnteryMoney(amount int) (*mongo.UpdateResult, error) {
	match := bson.M{"_id": receiver.Id}
	change := bson.M{"$set": bson.M{"enteredmoneies": &Enterdmoney{
		Amount: amount,
		Time:   time.Now(),
	}}}
	one, err := db.CountingCollection.UpdateOne(context.Background(), match, change)
	return one, err

}
