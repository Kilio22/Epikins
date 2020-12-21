package main

import (
	"context"
	"epikins-api/internal/services/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

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

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mongoClient := connectToDb()
	defer func() {
		err := mongoClient.Disconnect(context.TODO())
		log.Println(err)
	}()

	appData := initAppData(mongoClient)
	app := setupApp(appData)
	appPort := utils.GetEnvVariable("SERVER_PORT")
	log.Println("Listening on port " + appPort)
	log.Fatal(app.Listen(":" + appPort))
}
