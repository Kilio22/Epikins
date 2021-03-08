package main

import (
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/controllers/jenkinsCredentials"
	"epikins-api/internal/controllers/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func setupApp(appData *internal.AppData) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(func(ctx *fiber.Ctx) error {
		return isAuthenticated(appData, ctx)
	})

	app.Post("/login", func(ctx *fiber.Ctx) error {
		return controllers.LoginController(appData, ctx)
	})

	projectsGroup := app.Group("/projects", func(ctx *fiber.Ctx) error {
		return checkUserRole(appData, ctx, config.PROJECTS, config.MODULE)
	})
	projectsGroup.Get("/", func(ctx *fiber.Ctx) error {
		return controllers.ProjectsController(appData, ctx)
	})
	projectsGroup.Get("/:module/:project", func(ctx *fiber.Ctx) error {
		return controllers.ProjectInformationController(appData, ctx)
	})
	projectsGroup.Get("/:module/:project/:city", func(ctx *fiber.Ctx) error {
		return controllers.ProjectJobsController(appData, ctx)
	})

	buildGroup := app.Group("/build", func(ctx *fiber.Ctx) error {
		return checkUserRole(appData, ctx, config.PROJECTS)
	})
	buildGroup.Post("/", func(ctx *fiber.Ctx) error {
		return controllers.BuildController(appData, ctx)
	})
	buildGroup.Post("/global", func(ctx *fiber.Ctx) error {
		return controllers.GlobalBuildController(appData, ctx)
	})

	getCredentialsGroup := app.Group("/credentials", func(ctx *fiber.Ctx) error {
		return checkUserRole(appData, ctx, config.USERS, config.CREDENTIALS)
	})
	getCredentialsGroup.Get("/", func(ctx *fiber.Ctx) error {
		return jenkinsCredentials.GetJenkinsCredentialsController(appData, ctx)
	})

	protectedCredentialsGroup := app.Group("/credentials", func(ctx *fiber.Ctx) error {
		return checkUserRole(appData, ctx, config.CREDENTIALS)
	})
	protectedCredentialsGroup.Post("/", func(ctx *fiber.Ctx) error {
		return jenkinsCredentials.AddJenkinsCredentialController(appData, ctx)
	})
	protectedCredentialsGroup.Delete("/:username", func(ctx *fiber.Ctx) error {
		return jenkinsCredentials.DeleteJenkinsCredentialController(appData, ctx)
	})

	usersGroup := app.Group("/users", func(ctx *fiber.Ctx) error {
		return checkUserRole(appData, ctx, config.USERS)
	})
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

	moduleManagerGroup := app.Group("/projects", func(ctx *fiber.Ctx) error {
		return checkUserRole(appData, ctx, config.MODULE)
	})
	moduleManagerGroup.Put("/:module/:project", func(ctx *fiber.Ctx) error {
		return controllers.UpdateProjectBuildLimitController(appData, ctx)
	})

	studentGroup := app.Group("/student")
	studentGroup.Get("/jobs", func(ctx *fiber.Ctx) error {
		return controllers.StudentJobsController(appData, ctx)
	})
	studentGroup.Post("/build", func(ctx *fiber.Ctx) error {
		return controllers.StudentBuildController(appData, ctx)
	})

	logGroup := app.Group("/log", func(ctx *fiber.Ctx) error {
		return checkUserRole(appData, ctx, config.LOG)
	})
	logGroup.Get("/", func(ctx *fiber.Ctx) error {
		return controllers.GetBuildLogsController(appData, ctx)
	})

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})
	return app
}
