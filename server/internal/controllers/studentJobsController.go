package controllers

import (
	"net/http"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/controllers/util"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func StudentJobsController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	jenkinsCredentials, err := util.GetJenkinsCredentials(config.HighestPrivilegeJenkinsLogin, appData.JenkinsCredentialsCollection)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return SendMessage(c, "cannot get student jobs: no jenkins credentials found with login \""+config.HighestPrivilegeJenkinsLogin+"\"", http.StatusInternalServerError)
		}
		return SendMessage(c, "cannot get student jobs: "+err.Error(), http.StatusInternalServerError)
	}
}
