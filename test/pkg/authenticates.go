package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"radical.com/go-rest-api/db"
	"strconv"
)

var collection = db.UserCollection

func dehashuseruser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	fmt.Printf("claims %v\n\n", claims)
	admin := claims["admin"].(bool)
	println(admin)
	name := claims["user"].(map[string]interface{})
	bytes, e := json.Marshal(name)
	if e != nil {
		return c.SendString(e.Error())
	}
	return c.SendString(string(bytes))
}

func lovehandle(c *fiber.Ctx) error {
	get := c.Get("user-Agent")
	//fmt.Printf("%s", get)
	c.SendString(get)
	return nil
}

func Authentically(app *fiber.App, collection *mongo.Collection) {
	app.Get("/dehashuser", dehashuseruser)
	app.Get("/delallusers", delallusers)
	app.Get("/delallphones", deleteAllCodes)

}

func deleteAllCodes(c *fiber.Ctx) error {
	deleteMany, err := db.DpPhoneCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendString(strconv.Itoa(int(deleteMany.DeletedCount)))
}

func delallusers(c *fiber.Ctx) error {
	deleteMany, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendString(strconv.Itoa(int(deleteMany.DeletedCount)))
}
