package migrations

// import (
// 	"context"
// 	"crs/models"
// 	"fmt"

// 	"go.mongodb.org/mongo-driver/mongo"
// )

// // DB es la conexión a la base de datos
// var DB *mongo.Database

// // Up1 crea la colección Temperatures si no existe
// func Up1() error {
// 	// Crea la tabla Temperatures si no existe
// 	err := DB.AutoMigrate(&models.Temperature{})
// 	return err
// }

// // Down1 elimina la colección Temperatures
// func Down1() error {
// 	// Elimina la tabla Temperatures
// 	err := DB.Collection("Temperatures").Drop(context.Background())
// 	return err
// }

// func main() {
// 	// Importa el valor de DB desde database
// 	DB = database.InitDB()

// 	// Ejecuta la migración
// 	err := migrations.Up1()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Migración ejecutada con éxito")
// }
