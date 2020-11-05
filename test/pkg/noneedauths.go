package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math/rand"
	"radical.com/go-rest-api/db"
	"radical.com/go-rest-api/test/model"
	"time"
)

var Collection = db.UserCollection
var ctx = context.Background()

func NonAuthentically(app *fiber.App) {
	app.Get("/hate", lovehandle)
	app.Post("/login", login)
	app.Post("/dynamicpass", dynamicpass)
	app.Post("/dp/verify", dpverify)

}

func login(c *fiber.Ctx) error {
	username := c.FormValue("user")
	pass := c.FormValue("pass")

	fmt.Printf("%v, %v\n", username, pass)

	//verify is username and pass is in db and corrected
	var user model.User
	if err := usercollection.FindOne(ctx, bson.M{"username": username}).Decode(&user); err != nil {
		//log.Fatal(err)
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(map[string]interface{}{"status": "fail",
			"code": "user not found"})
	}

	if db.ComparePasswords(user.Password, pass) {
		fmt.Printf("%v", user)

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)
		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		//asin := math.Asin(rand.Float64())
		//us := &model.User{
		//	FirstName: "god" + strconv.FormatFloat(asin, 'f', 6, 64),
		//	LastName:  "king",
		//	Username:  username,
		//}
		claims["user"] = user
		claims["admin"] = false
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		fmt.Printf("%v", claims)
		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	}
	return c.SendStatus(fiber.StatusUnauthorized)
}

func dpverify(c *fiber.Ctx) error {
	phoneNumber := c.FormValue("phonenumber")
	code := c.FormValue("code")
	tokenmap, exception := verifyandtoken(&phoneNumber, &code)
	if exception != nil {
		return exception
	}
	return c.JSON(tokenmap)
}

func verifyandtoken(number, code *string) (fiber.Map, model.VerificationCOdeException) {
	_, exception := model.Verifycoderecieved(number, code)
	if exception != nil {
		return nil, exception
	} else {
		//add user with at least data to db
		insertOneResult, err := model.InsertNewUserByPhoneNumber(*number)
		id := insertOneResult.InsertedID.(primitive.ObjectID)

		token, err := generateTokenbasephone(id, *number)
		return token, err
	}
}

func generateTokenbasephone(number primitive.ObjectID, phonenumber string) (fiber.Map, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = model.User{Id: number, Phonenumber: phonenumber}
	claims["role"] = 0
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	fmt.Printf("new claims %v\n", claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}
	return fiber.Map{"token": t}, nil
}
func generateTokenbaseInfoMap(infomap map[string]interface{}) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = infomap
	claims["role"] = 0
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	fmt.Printf("new claims %v\n", claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func dynamicpass(c *fiber.Ctx) error {
	phoneNumber := c.FormValue("phonenumber")
	if checkPhoneNUmberValid(phoneNumber) {
		addtodpverifying(phoneNumber)
		return c.SendString("{\"status\":\"ok\"}")
	} else {
		return errors.New("no valid number")
	}
}

func addtodpverifying(phoneNumber string) {
	code := ""
	for i := 0; i < 5; i++ {
		code += (string)(rune(rand.Intn(10) + 48))
	}
	model.Inserttodpphonedb(phoneNumber, code)
}

func checkPhoneNUmberValid(number string) bool {
	return true
}
