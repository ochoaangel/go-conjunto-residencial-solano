package main

import (
	"crs/controllers"
	"crs/initializers"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database

func main() {

	err := myInit()

	if err != nil {
		panic(err)
	}

}

func myInit() error {

	// init env
	err := initializers.LoadEnvVariables()
	if err != nil {
		return err
	}

	// init DB
	err = initializers.InitDB()
	if err != nil {
		return err
	}

	// defer closing db
	defer initializers.CloseDB()

	// start server
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	// create App
	app := gin.Default()

	app.POST("/", controllers.Create)
	app.GET("/getAll", controllers.GetAll)
	app.GET("/getRange", controllers.GetRange)
	app.GET("/getLastNotNull", controllers.GetLastNotNull)
	app.GET("/getLastWithNull", controllers.GetLastWithNull)
	app.GET("/test", controllers.Test)

	// Correr App
	app.Run()

	return nil
}
