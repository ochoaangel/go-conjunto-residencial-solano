package controllers

import (
	"context"
	"crs/initializers"
	"crs/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

func TemperatureCreateOne(c *gin.Context) {
	collection := initializers.DB.Collection("Temperatura")
	var body models.Temperature

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "Temperature NOT added",
			"error":   err.Error(),
		})
		return
	}

	body.CreatedAt = time.Now()

	result, err := collection.InsertOne(context.Background(), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Temperature NOT added",
			"error":   "Failed to create temperature",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Temperature added",
		"newId":   result.InsertedID,
	})
}

func TemperatureCreateMultiple(c *gin.Context) {
	collection := initializers.DB.Collection("Temperatura")

	var bodies []models.Temperature

	if err := c.BindJSON(&bodies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	validate := validator.New()
	for _, body := range bodies {
		err := validate.Struct(body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"ok":      false,
				"message": "Temperatures NOT added",
				"error":   err.Error(),
			})
			return
		}

		body.CreatedAt = time.Now()
	}

	interfaces := convertTemperaturesToInterfaces(bodies)

	result, err := collection.InsertMany(context.Background(), interfaces)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Temperatures NOT added",
			"error":   "Failed to create temperatures",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Temperatures added",
		"newIds":  result.InsertedIDs,
	})
}

func convertTemperaturesToInterfaces(temperatures []models.Temperature) []interface{} {
	var interfaces []interface{}

	for _, temperature := range temperatures {
		interfaces = append(interfaces, bson.M{
			"Temp1": temperature.Temp1,
			"Temp2": temperature.Temp2,
			"Temp3": temperature.Temp3,
		})
	}

	return interfaces
}

func TemperatureGetOneByID(c *gin.Context) {
	collection := initializers.DB.Collection("Temperatura")
	id := c.Param("id")

	var result models.Temperature
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"ok":      false,
			"message": "Temperature not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Temperature found",
		"data":    result,
	})
}

func TemperatureGetMultipleByID(c *gin.Context) {
	collection := initializers.DB.Collection("Temperatura")
	ids := c.QueryArray("ids")

	var results []models.Temperature
	cursor, err := collection.Find(context.Background(), bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Failed to fetch temperatures",
			"error":   err.Error(),
		})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var temp models.Temperature
		err := cursor.Decode(&temp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ok":      false,
				"message": "Failed to fetch temperatures",
				"error":   err.Error(),
			})
			return
		}
		results = append(results, temp)
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Temperatures found",
		"data":    results,
	})
}

func TemperatureGetMultipleByCreatedAtRange(c *gin.Context) {
	collection := initializers.DB.Collection("Temperatura")
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	var results []models.Temperature
	cursor, err := collection.Find(context.Background(), bson.M{"created_at": bson.M{"$gte": startDate, "$lte": endDate}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Failed to fetch temperatures",
			"error":   err.Error(),
		})
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var temp models.Temperature
		err := cursor.Decode(&temp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"ok":      false,
				"message": "Failed to fetch temperatures",
				"error":   err.Error(),
			})
			return
		}
		results = append(results, temp)
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Temperatures found",
		"data":    results,
	})
}

func TemperatureDeleteOneByID(c *gin.Context) {
	collection := initializers.DB.Collection("Temperatura")
	id := c.Param("id")

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Failed to delete temperature",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Temperature deleted",
		"count":   result.DeletedCount,
	})
}

func TemperatureDeleteMultipleByID(c *gin.Context) {
	collection := initializers.DB.Collection("Temperatura")
	ids := c.QueryArray("ids")

	result, err := collection.DeleteMany(context.Background(), bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ok":      false,
			"message": "Failed to delete temperatures",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Temperatures deleted",
		"count":   result.DeletedCount,
	})
}

// func TemperatureCreate(c *gin.Context) {
// 	collection := initializers.DB.Collection("Temperatura")

// 	// Estructura esperada en el cuerpo
// 	var body models.Temperature

// 	// Validar los datos de entrada
// 	if err := c.BindJSON(&body); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Invalid request body",
// 		})
// 		return
// 	}

// 	// Validar la estructura de datos
// 	validate := validator.New()
// 	err := validate.Struct(body)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"ok":      false,
// 			"message": "Temperatures NOT added",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}

// 	// Aggrega el valoractual
// 	body.CreatedAt = time.Now()

// 	// Crear el modelo de estructura sin agregar elementos
// 	model := models.Temperature{
// 		Temp1:     body.Temp1,
// 		Temp2:     body.Temp2,
// 		Temp3:     body.Temp3,
// 		CreatedAt: body.CreatedAt,
// 	}

// 	result, err := collection.InsertOne(context.Background(), model)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"ok":      false,
// 			"message": "Temperatures NOT added",
// 			"error":   "Failed to create temperature",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"ok":      true,
// 		"message": "Temperatures added",
// 		"newId":   result.InsertedID,
// 	})
// }
