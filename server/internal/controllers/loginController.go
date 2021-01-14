package controllers

import (
	"context"

	"epikins-api/config"
	"epikins-api/internal"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
)

func LoginController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	var user internal.User

	err := appData.UsersCollection.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		return c.JSON(internal.User{
			Email: userEmail,
			Roles: []internal.Role{config.STUDENT},
		})
	}
	user.Roles = append(user.Roles, config.STUDENT)
	return c.JSON(user)
}
