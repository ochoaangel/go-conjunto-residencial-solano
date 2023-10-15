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

	app.POST("/Temperature/createOne", controllers.TemperatureCreateOne)
	app.POST("/Temperature/createMultiple", controllers.TemperatureCreateMultiple)
	app.GET("/Temperature/getOne/:id", controllers.TemperatureGetOneByID)
	app.GET("/Temperature/getMultiple", controllers.TemperatureGetMultipleByID)
	app.GET("/Temperature/getByCreatedAtRange", controllers.TemperatureGetMultipleByCreatedAtRange)
	app.DELETE("/Temperature/deleteOne/:id", controllers.TemperatureDeleteOneByID)
	app.DELETE("/Temperature/deleteMultiple", controllers.TemperatureDeleteMultipleByID)

	// Correr App
	app.Run()

	return nil
}
