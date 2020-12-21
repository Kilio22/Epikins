package main

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/controllers/jenkinsCredentials"
	"epikins-api/internal/controllers/users"
	"epikins-api/internal/services/loginService"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

func isAuthenticated(appData *internal.AppData, c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")
	email, err := loginService.LoginService(appData.AppId, accessToken)
	if err != nil {
		return controllers.SendMessage(c, err.Error(), http.StatusUnauthorized)
	}
	c.Request().Header.Set("email", email)
	return c.Next()
}

func setupApp(appData *internal.AppData) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(func(ctx *fiber.Ctx) error {
		return isAuthenticated(appData, ctx)
	})

	app.Post("/login", controllers.LoginController)
	app.Get("/projects", func(ctx *fiber.Ctx) error {
		return controllers.ProjectsController(appData, ctx)
	})
	app.Get("/projects/:project", func(ctx *fiber.Ctx) error {
		return controllers.ProjectJobsController(appData, ctx)
	})

	buildGroup := app.Group("/build")
	buildGroup.Post("/", func(ctx *fiber.Ctx) error {
		return controllers.BuildController(appData, ctx)
	})
	buildGroup.Post("/global", func(ctx *fiber.Ctx) error {
		return controllers.GlobalBuildController(appData, ctx)
	})

	credentialsGroup := app.Group("/credentials")
	credentialsGroup.Get("/", func(ctx *fiber.Ctx) error {
		return jenkinsCredentials.GetJenkinsCredentialsController(appData, ctx)
	})
	credentialsGroup.Post("/", func(ctx *fiber.Ctx) error {
		return jenkinsCredentials.AddJenkinsCredentialController(appData, ctx)
	})
	credentialsGroup.Put("/", func(ctx *fiber.Ctx) error {
		return jenkinsCredentials.UpdateJenkinsCredentialsController(appData, ctx)
	})
	credentialsGroup.Delete("/:username", func(ctx *fiber.Ctx) error {
		return jenkinsCredentials.DeleteJenkinsCredentialController(appData, ctx)
	})

	usersGroup := app.Group("/users")
	usersGroup.Get("/", func(ctx *fiber.Ctx) error {
		return users.GetUsersController(appData, ctx)
	})
	usersGroup.Post("/", func(ctx *fiber.Ctx) error {
		return users.AddUserController(appData, ctx)
	})
	usersGroup.Put("/", func(ctx *fiber.Ctx) error {
		return users.UpdateUserController(appData, ctx)
	})
	usersGroup.Delete("/:username", func(ctx *fiber.Ctx) error {
		return users.DeleteUserController(appData, ctx)
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	return app
}
