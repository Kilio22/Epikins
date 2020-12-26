package controllers

import (
	"context"
	"epikins-api/internal"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoginController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	var user internal.User

	err := appData.UsersCollection.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		return SendMessage(c, "something went wrong: "+err.Error(), http.StatusInternalServerError)
	}
	return c.JSON(user)
}
