package main

import (
	"context"
	"log"

	"epikins-api/config"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDb() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://" + util.GetEnvVariable(config.MongoHostKey) + ":" + util.GetEnvVariable(config.MongoPortKey))
	clientOptions.SetAuth(options.Credential{
		Username: util.GetEnvVariable(config.MongoUsernameKey), Password: util.GetEnvVariable(config.MongoPasswordKey),
	})

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
	appPort := util.GetEnvVariable(config.ServerPortKey)
	log.Println("Listening on port " + appPort)
	log.Fatal(app.Listen(":" + appPort))
}
