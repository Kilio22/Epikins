package main

import (
	"context"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func hasRole(roles []internal.Role, acceptedRoles []internal.Role) bool {
	for _, role := range roles {
		for _, acceptedRole := range acceptedRoles {
			if role == acceptedRole {
				return true
			}
		}
	}
	return false
}

func checkUserRole(appData *internal.AppData, c *fiber.Ctx, acceptedRoles ...internal.Role) error {
	var user internal.User
	err := appData.UsersCollection.FindOne(context.TODO(), bson.M{"email": c.Get("email")}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return controllerUtil.SendMyError(util.GetMyError("you're not authorized to access to this resource", nil, http.StatusForbidden), c)
		}
		return controllerUtil.SendMyError(util.GetMyError("cannot check user role", err, http.StatusInternalServerError), c)
	}
	if hasRole(user.Roles, acceptedRoles) == false {
		return controllerUtil.SendMyError(util.GetMyError("you're not authorized to access to this resource", nil, http.StatusForbidden), c)
	}
	return c.Next()
}
