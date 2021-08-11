package main

import (
	"context"
	"fmt"
	"log"
	"study_car_sharing/api"
	. "study_car_sharing/config"

	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config, err := GetConfig()
	if err != nil {
		panic(err)
	}
	fmt.Print(config)
	fmt.Print("start server")
	startServer(config)
}
func startServer(config Config) error {
	//client, conn := DatabaseConnection(config.Database)
	//defer DatabaseDisconnect(context.TODO(), client)
	router := api.ServerRunner{}

	//log.Printf(conn.Name())
	log.Print(http.ListenAndServe(fmt.Sprintf(":%d", 5001), router.Route()))
	//mongo:=conn.Database(config.Database.DBName).Collection(config.Database.CollectionName)
	return nil
}
func DatabaseConnection(dbConfig DatabaseConfig) (*mongo.Client, *mongo.Database) {
	clientOptions := options.Client().ApplyURI("mongodb://" + dbConfig.Host + ":" + dbConfig.Port).
		SetAuth(options.Credential{
			Username: dbConfig.User,
			Password: dbConfig.Password,
		})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Made")
	mongo := client.Database(dbConfig.DBName)
	return client, mongo
}

func DatabaseDisconnect(ctx context.Context, client *mongo.Client) {
	client.Disconnect(ctx)
}
