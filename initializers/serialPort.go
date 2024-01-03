package initializers

import (
	"bufio"
	"crs/models"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/jacobsa/go-serial/serial"
)

var Port io.ReadWriteCloser

func SerialPortConfig(crs models.CrsInitConfigModel) {

	// puerto por defecto
	com := crs.Comunicacion.Port
	if com == "" {
		com = "COM5"
	}

	// baud o velocidad por defecto
	baud := crs.Comunicacion.Baud
	if baud == 0 {
		baud = 9606
	}

	options := serial.OpenOptions{
		PortName:        com,
		BaudRate:        uint(baud),
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	go func() {
		port, err := serial.Open(options)
		if err != nil {
			fmt.Println("ERROR >>>>>>>>>> No se pudo abrir el puerto SERIE, intentando nuevamente en 1 minuto con:", com, " a ", baud)
			for {
				time.Sleep(1 * time.Second)
				port, err = serial.Open(options)
				if err == nil {
					fmt.Println(">>>>>>>>>> Puerto SERIE abierto con éxito:", com, " a ", baud)
					break
				}
			}
		} else {
			fmt.Println(">>>>>>>>>> Puerto SERIE abierto con éxito:", com, " a ", baud)

		}

		reader := bufio.NewReader(port)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("ERROR >>>>>>>>>> Se perdió la conexión con el puerto, intentando reconectar...", com, " a ", baud)
				for {
					time.Sleep(1 * time.Second)
					port, err = serial.Open(options)
					if err == nil {
						fmt.Println(">>>>>>>>>> Puerto SERIE abierto con éxito:", com, " a ", baud)
						reader = bufio.NewReader(port)
						break
					}
				}
			} else {
				message = strings.TrimSpace(message)
				fmt.Printf("Mensaje recibido por puerto Serial: %s\n", message)
			}
		}
	}()

}
