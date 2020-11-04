package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"radical.com/go-rest-api/db"
)

type Phone struct {
	PhoneNumber string `json:"phonenumber"`
	Code        string `json:"code"`
}

type VerificationCOdeException interface {
	error
}

func Verifycodephone(number, code *string) (bool bool, verificationCOdeException VerificationCOdeException) {
	filter := bson.M{"phonenumber": number}
	result := db.DpPhoneCollection.FindOne(context.Background(), filter)
	phoneobj := Phone{}
	er := result.Decode(&phoneobj)
	if er != nil {
		return false, er
	}
	if phoneobj.Code == *code {
		return true, nil
	} else
	{
		return false, nil
	}

}

func Inserttodpphonedb(phonenumber string, code string) bool {
	upsert := true
	pattern := "^" + phonenumber + "$|" + "^$"
	println(pattern)
	filter := bson.M{"phonenumber": primitive.Regex{Pattern: pattern, Options: ""}}
	result := db.DpPhoneCollection.FindOne(context.Background(), filter)
	phone := Phone{}
	_ = result.Decode(&phone)
	//fmt.Printf("result  number is %v\n", phone)
	//fmt.Printf("phone number is %s\n", phonenumber)
	updateM := bson.M{"$set": bson.M{"phonenumber": phonenumber, "code": code}}
	_, err := db.DpPhoneCollection.UpdateOne(context.Background(),
		filter,
		updateM,
		&options.UpdateOptions{Upsert: &upsert})
	fmt.Printf("%v", err)
	return err != nil
}
