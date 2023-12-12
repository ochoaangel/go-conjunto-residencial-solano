package controllers

import (
	"context"
	"crs/initializers"
	"crs/models"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAll(c *gin.Context) {
	collection := initializers.DB.Collection("All")
	var elements []models.All

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener las variables",
		})
		return
	}

	if err = cursor.All(context.Background(), &elements); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener las variables",
		})
		return
	}

	c.JSON(http.StatusOK, elements)
}

func GetRange(c *gin.Context) {
	collection := initializers.DB.Collection("All")
	var elements []models.All
	var rangeData models.Range

	if err := c.BindJSON(&rangeData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	from, _ := time.Parse(time.RFC3339, rangeData.From)
	to, _ := time.Parse(time.RFC3339, rangeData.To)

	if from.After(to) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "La fecha 'from' debe ser menor que la fecha 'to'",
		})
		return
	}

	cursor, err := collection.Find(context.Background(), bson.M{
		"created_at": bson.M{
			"$gte": from,
			"$lte": to,
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener los valores entre las fechas",
		})
		return
	}

	if err = cursor.All(context.Background(), &elements); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener los valores entre las fechas",
		})
		return
	}

	c.JSON(http.StatusOK, elements)
}

func GetLastNotNull(c *gin.Context) {
	collection := initializers.DB.Collection("All")
	var element models.All

	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "fecha", Value: -1}})
	err := collection.FindOne(context.Background(), bson.M{}, opts).Decode(&element)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener los últimos valores válidos",
		})
		return
	}

	c.JSON(http.StatusOK, element)
}

func GetLastWithNull(c *gin.Context) {
	collection := initializers.DB.Collection("All")
	var element models.All

	opts := options.FindOne().SetSort(bson.D{primitive.E{Key: "fecha", Value: -1}})
	err := collection.FindOne(context.Background(), bson.M{}, opts).Decode(&element)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener el elemento más reciente",
		})
		return
	}

	c.JSON(http.StatusOK, element)
}

func Create(c *gin.Context) {
	collection := initializers.DB.Collection("All")
	var body models.All

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
			"error": err.Error(),
		})
		return
	}

	// Verificar si al menos una clave válida está presente
	v := reflect.ValueOf(body)
	typeOfS := v.Type()

	isValid := false
	for i := 0; i < v.NumField(); i++ {
		if typeOfS.Field(i).Name != "ID" && typeOfS.Field(i).Name != "CreatedAt" && !reflect.ValueOf(v.Field(i).Interface()).IsZero() {
			isValid = true
			break
		}
	}

	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Debe haber al menos una clave válida",
		})
		return
	}

	body.CreatedAt = time.Now()

	// Generar un nuevo valor de campo `_id` automáticamente.
	body.ID = primitive.NewObjectID()

	result, err := collection.InsertOne(context.Background(), body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al crear la variable",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"newId": result.InsertedID,
	})
}

func Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Conexión correcta",
	})
}
