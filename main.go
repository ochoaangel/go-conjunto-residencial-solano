package main

import (
	"crs/controllers"
	"crs/initializers"
	"crs/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database
var crs models.CrsInitConfigModel

func main() {

	///////////////////////////////////////////////////////////////////////////////////////////
	// Verifico si existe otra intancia de este progrma abierto
	///////////////////////////////////////////////////////////////////////////////////////////

	// verifico si existe el archivo lockfile.lock
	content, err := ioutil.ReadFile("lockfile.lock")
	if err == nil {
		pid, err := strconv.Atoi(string(content))
		if err != nil {
			fmt.Println("Error al leer el PID del archivo de bloqueo:", err)
			os.Exit(1)
		}

		// Verificar si el proceso está en ejecución
		_, err = os.FindProcess(pid)
		if err == nil {
			fmt.Println("El programa ya está en ejecución, saliendo...")
			os.Exit(1)
		} else {
			fmt.Println("El programa anterior no está en ejecución, eliminando el archivo de bloqueo...")
			os.Remove("lockfile.lock")
		}

	}

	// obtengo el PID
	pid := os.Getpid()

	// guardo el PID en el fichero lockfile.lock
	err2 := ioutil.WriteFile("lockfile.lock", []byte(strconv.Itoa(pid)), 0644)
	if err2 != nil {
		fmt.Println("Error al crear el archivo de bloqueo:", err2)
		return
	}

	// en caso de cerrar la aplicación
	defer os.Remove("lockfile.lock")

	fmt.Println("El programa se está ejecutando.")

	// Manejar la señal de interrupción para limpiar antes de salir (trata de eliminar rastros al cerrar inesperadamente)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove("lockfile.lock")
		os.Exit(1)
	}()

	///////////////////////////////////////////////////////////////////////////////////////////
	// Verifico si existe el archivo 'crs-init-config.json' de configuraciones iniciales
	///////////////////////////////////////////////////////////////////////////////////////////

	// Leer el archivo JSON de configuración
	data, err := ioutil.ReadFile("crs-init-config.json")
	if err != nil {
		log.Fatalf("Error al leer el archivo 'crs-init-config.json': %v", err)
	}

	// Deserializar el JSON en la estructura Config
	err = json.Unmarshal(data, &crs)
	if err != nil {
		log.Fatalf("Error al deserializar el JSON 'crs-init-config.json' por: %v", err)
	}

	///////////////////////////////////////////////////////////////////////////////////////////
	// Inicializo mi programa
	///////////////////////////////////////////////////////////////////////////////////////////
	err3 := myInit()

	if err3 != nil {
		panic(err)
	}

}

func myInit() error {

	// init env
	err := initializers.LoadEnvVariables()
	if err != nil {
		return err
	}

	// init Serial
	initializers.SerialPortConfig(crs)

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
