package main

import (
	"context"
	"epikins-api/internal/services/utils"
	"log"
	"strconv"

	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/pkg/libJenkins"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ProjectsCollectionName string = "projects"

func connectToDb() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://" + utils.GetEnvVariable("MONGO_HOST") + ":" + utils.GetEnvVariable("MONGO_PORT"))
	clientOptions.SetAuth(options.Credential{Username: utils.GetEnvVariable("MONGO_INITDB_ROOT_USERNAME"), Password: utils.GetEnvVariable("MONGO_INITDB_ROOT_PASSWORD")})

	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient
}

func setupApp(collection *mongo.Collection) *fiber.App {
	app := fiber.New()
	appData := &internal.AppData{Collection: collection, ProjectsData: make(map[libJenkins.AccountType]internal.ProjectsData), AppId: utils.GetEnvVariable("APP_ID")}

	app.Use(cors.New())
	app.Post("/login", func(ctx *fiber.Ctx) {
		controllers.LoginController(appData, ctx)
	})
	app.Get("/projects", func(ctx *fiber.Ctx) {
		controllers.ProjectsController(appData, ctx)
	})
	app.Get("/projects/:project", func(ctx *fiber.Ctx) {
		controllers.ProjectJobsController(appData, ctx)
	})
	app.Post("/build", func(ctx *fiber.Ctx) {
		controllers.BuildController(appData, ctx)
	})
	app.Post("/build/global", func(ctx *fiber.Ctx) {
		controllers.GlobalBuildController(appData, ctx)
	})
	app.Use(func(c *fiber.Ctx) {
		c.SendStatus(404)
	})
	return app
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mongoClient := connectToDb()
	epikinsDatabase := mongoClient.Database(utils.GetEnvVariable("MONGO_DB"))
	projectsCollection := epikinsDatabase.Collection(ProjectsCollectionName)
	defer func() {
		err := mongoClient.Disconnect(context.TODO())
		log.Println(err)
	}()

	app := setupApp(projectsCollection)
	appPort, err := strconv.Atoi(utils.GetEnvVariable("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Listening on port %d\n", appPort)
	err = app.Listen(appPort)
	if err != nil {
		log.Fatal(err)
	}
}
