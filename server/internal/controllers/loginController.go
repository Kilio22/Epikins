package controllers

import (
	"context"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
)

func LoginController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	var user internal.User

	err := appData.UsersCollection.FindOne(context.TODO(), bson.M{"email": userEmail}).Decode(&user)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError("cannot log user in", err, http.StatusInternalServerError), c)
	}
	return c.JSON(user)
}
