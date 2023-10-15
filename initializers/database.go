package initializers

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return errors.New("you must set your 'MONGODB_URI' environmental variable")
	}

	dbName := os.Getenv("MONGODB_NAME")
	if dbName == "" {
		return errors.New("you must set your 'MONGODB_NAME' environmental variable")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	DB = client.Database(dbName)

	// // -------------------------------------------------------------------
	// collection := DB.Collection("Temperatura")

	// // Crear el modelo de estructura sin agregar elementos
	// model := models.Temperature{
	// 	Temp1:     25.5,
	// 	Temp2:     26.0,
	// 	Temp3:     24.8,
	// 	CreatedAt: time.Now(),
	// }

	// _, err = collection.InsertOne(context.Background(), model)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Modelo de estructura agregado a la colección")

	// // Verificar si la base de datos existe
	// err = DB.RunCommand(context.Background(), map[string]interface{}{"ping": 1}).Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Crear la colección "Temperatura" con el modelo si no existe
	// err = DB.CreateCollection(context.Background(), "Temperatura")
	// if err != nil && !mongo.IsDuplicateKeyError(err) {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Base de datos y colección creadas")

	return nil
}

func CloseDB() error {
	return DB.Client().Disconnect(context.Background())
}
