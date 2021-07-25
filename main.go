package main

import (
	"context"
	"fmt"
	"log"
	. "study_car_sharing/config"

	"net/http"

	"github.com/gorilla/mux"
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
	conn := DatabaseConnection(config.Database)
	router := mux.NewRouter()
	log.Printf(conn.Name())
	log.Print(http.ListenAndServe(fmt.Sprintf(":%d", 5000), router))
	//mongo:=conn.Database(config.Database.DBName).Collection(config.Database.CollectionName)
	return nil
}
func DatabaseConnection(dbConfig DatabaseConfig) *mongo.Database {
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
	return mongo
}
