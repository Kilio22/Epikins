package main

import (
	"context"
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func hasRole(roles []internal.Role, toFind internal.Role) bool {
	for _, role := range roles {
		if role == toFind {
			return true
		}
	}
	return false
}

func checkUserRole(appData *internal.AppData, role internal.Role, c *fiber.Ctx) error {
	var user internal.User
	err := appData.UsersCollection.FindOne(context.TODO(), bson.M{"email": c.Get("email")}).Decode(&user)

	if err != nil {
		return err
	}
	if hasRole(user.Roles, role) == false {
		return controllers.SendMessage(c, "you're not authorized to access to this resource", http.StatusForbidden)
	}
	return c.Next()
}
