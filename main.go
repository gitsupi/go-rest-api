package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"radical.com/go-rest-api/db"
	"reflect"
)

const dbName = "personsdb"
const collectionName = "person"
const port = 8000

var collection, err = db.GetMongoDbCollection(dbName, collectionName)

func getPerson(c *fiber.Ctx) error {
	//collection, err := db.GetMongoDbCollection(dbName, collectionName)
	if err != nil {
		err, _ := json.Marshal(err)
		return c.Status(500).Send(err)

	}

	var filter = bson.M{}
	if c.Params("id") != "" {
		id := c.Params("id")
		objID, _ := primitive.ObjectIDFromHex(id)
		filter = bson.M{"_id": objID}
	}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())

	if err != nil {
		err, _ := json.Marshal(err)
		return c.Status(500).Send(err)
	}

	cur.All(context.Background(), &results)

	if results == nil {
		return c.SendStatus(404)
	}

	j, _ := json.Marshal(results)

	return c.Send(j)
}

func getAllPersons(c *fiber.Ctx) error {
	createPerson(c)
	if err != nil {
		err, _ := json.Marshal(err)
		return c.Status(500).Send(err)

	}

	var filter bson.M = bson.M{}
	//
	//if c.Params("id") != "" {
	//	id := c.Params("id")
	//	objID, _ := primitive.ObjectIDFromHex(id)
	//	filter = bson.M{"_id": objID}
	//}

	var results []bson.M
	cur, err := collection.Find(context.Background(), filter)
	defer cur.Close(context.Background())
	if err != nil {
		err, _ := json.Marshal(err)
		return c.Status(500).Send(err)

	}

	err = cur.All(context.Background(), &results)

	if results == nil || err != nil {
		return c.SendStatus(404)
	}

	reso, _ := json.Marshal(results)
	return c.Send(reso)
}

func createPerson(c *fiber.Ctx) error {
	//collection, err := db.GetMongoDbCollection(dbName, collectionName)
	if err != nil {
		err, _ := json.Marshal(err)
		c.Status(500).Send(err)
		return nil
	}

	var person db.Person
	_ = json.Unmarshal(c.Body(), &person)

	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		err, _ := json.Marshal(err)
		c.Status(500).Send(err)
		return nil
	}

	response, _ := json.Marshal(res)
	return c.Send(response)

}

func updatePerson(c *fiber.Ctx) error {
	var person db.Person
	unmarshaleEroor := json.Unmarshal([]byte(c.Body()), &person)
	if unmarshaleEroor != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(unmarshaleEroor.Error())
	}

	update := bson.M{
		"$set": person,
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)

	if err != nil {
		err, _ := json.Marshal(err)
		return c.Status(500).Send(err)
	}

	response, _ := json.Marshal(res)
	return c.Send(response)
}

func deletePerson(c *fiber.Ctx) error {
	//collection, err := db.GetMongoDbCollection(dbName, collectionName)

	if err != nil {
		err, _ := json.Marshal(err)

		c.Status(500).Send(err)
		return nil
	}

	objID, _ := primitive.ObjectIDFromHex(c.Params("id"))
	res, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})

	if err != nil {
		err, _ := json.Marshal(err)
		c.Status(500).Send(err)
		return nil
	}

	jsonResponse, _ := json.Marshal(res)
	return c.Send(jsonResponse)
}

func main() {
	var invs []int
	invs = append(invs, 12, 23, 23)
	history := db.History{HistoryInvestings: invs}
	print(reflect.TypeOf(history))

	person := db.Person{
		FirstName: "king",
		LastName:  "ming",
		Username:  "ming",
		Email:     "nf@nf.com",
		Counting: db.Counting{
			InvestedMoney: 12,
			ProfitedMoney: 0,
			History:       db.History{},
			InvestingType: 0,
		},
		Age: 0,
	}
	fmt.Printf("$%v", person)

	app := fiber.New()
	fmt.Printf("%v", history)
	app.Get("/person/:id?", getPerson)
	app.Post("/person", createPerson)
	app.Get("/persons", getAllPersons)
	app.Put("/person/:id", updatePerson)
	app.Delete("/person/:id", deletePerson)

	app.Listen(":1211")
}
