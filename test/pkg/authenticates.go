package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"radical.com/go-rest-api/db"
	"radical.com/go-rest-api/test/model"
	"strconv"
)

var usercollection = db.UserCollection

func Authentically(app *fiber.App, collection *mongo.Collection) {
	app.Get("/dehashuser", dehashuseruser)
	app.Get("/delallusers", delallusers)
	app.Get("/delallphones", deleteAllCodes)
	app.Post("/adduserinfo", adduserinfo)

}

func adduserinfo(c *fiber.Ctx) error {
	hashinfo := currentUser(c)
	user, err := model.GetUserByPhoneNumber(hashinfo["phonenumber"].(string))
	fmt.Printf("user %v\n", user)

	if err != nil {
		c.Status(401)
	}
	var out map[string]string
	err = c.BodyParser(&out)
	fmt.Printf("out %v\n", out)
	fmt.Printf("er %v\n", err)

	err = user.UpdateUserInfo(user, out)
	fmt.Printf("er2 %v\n", err)
	return c.Status(200).JSON(model.Status{
		Status:      "ok",
		Code:        1,
		Description: "ok is clear",
	})

}

func currentUser(c *fiber.Ctx) map[string]interface{} {
	var token *jwt.Token = c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	//fmt.Printf("\nclaims %v\n", claims)

	user, ok := claims["user"].(map[string]interface{})
	if !ok {
		c.Status(400)
	}
	return user
}

func dehashuseruser(c *fiber.Ctx) error {
	var token *jwt.Token = c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	fmt.Printf("\nclaims %v\n", claims)

	user, ok := claims["user"].(map[string]interface{})
	if ok {
		bytes, e := json.Marshal(user)
		if e != nil {
			return c.SendString(e.Error())
		}
		return c.SendString(string(bytes))
	}

	return c.JSON(model.Status{
		Status:      "nok",
		Code:        8888,
		Description: "can not convert user hash",
	})

}

func lovehandle(c *fiber.Ctx) error {
	get := c.Get("user-Agent")
	//fmt.Printf("%s", get)
	c.SendString(get)
	return nil
}

func deleteAllCodes(c *fiber.Ctx) error {
	deleteMany, err := db.DpPhoneCollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendString(strconv.Itoa(int(deleteMany.DeletedCount)))
}

func delallusers(c *fiber.Ctx) error {
	deleteMany, err := usercollection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return c.SendStatus(500)
	}
	return c.SendString(strconv.Itoa(int(deleteMany.DeletedCount)))
}
