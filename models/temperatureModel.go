package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Temperature struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Temp1     *float64           `bson:"temp1,omitempty"`
	Temp2     *float64           `bson:"temp2,omitempty"`
	Temp3     *float64           `bson:"temp3,omitempty"`
	CreatedAt time.Time          `bson:"created_at,	omitempty"`
}

// "*" se utiliza para declarar punteros a tipos de dato, hacee los campos opcional como estructura inicial.
// omitempty se utiliza para indicar que un campo debe ser omitido si su valor es el valor cero del tipo de dato correspondiente, el campo no se vería en la tabla si no tiene valor.
// validate:"required" hace que sea requerido
