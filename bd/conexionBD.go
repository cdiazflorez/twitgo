package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoC = conectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://cdiazflorez:Homero2022@cluster0.koxh1.mongodb.net/twitgo?retryWrites=true&w=majority")

/** Permite conectar a la BD de Mongo**/
func conectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa")
	return client
}

func CheckConnection() int {
	err := MongoC.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
