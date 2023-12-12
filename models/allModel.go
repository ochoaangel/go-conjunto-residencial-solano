package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type All struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	B1Temp    *float64           `bson:"b1temp,omitempty"`
	B2Temp    *float64           `bson:"b2temp,omitempty"`
	B3Temp    *float64           `bson:"b3temp,omitempty"`
	B1Term    *bool              `bson:"b1term,omitempty"`
	B2Term    *bool              `bson:"b2term,omitempty"`
	B3Term    *bool              `bson:"b3term,omitempty"`
	B1On      *bool              `bson:"b1on,omitempty"`
	B2On      *bool              `bson:"b2on,omitempty"`
	B3On      *bool              `bson:"b3on,omitempty"`
	DUltra    *float64           `bson:"dultra,omitempty"`
	PDigital  *float64           `bson:"pdigital,omitempty"`
	P1        *bool              `bson:"p1,omitempty"`
	P2        *bool              `bson:"p2,omitempty"`
	P3        *bool              `bson:"p3,omitempty"`
	P4        *bool              `bson:"p4,omitempty"`
	FCalle    *float64           `bson:"fcalle,omitempty"`
	Fin       *float64           `bson:"fin,omitempty"`
	NivelA1   *bool              `bson:"nivela1,omitempty"`
	NivelA2   *bool              `bson:"nivela2,omitempty"`
	NivelA3   *bool              `bson:"nivela3,omitempty"`
	NivelA4   *bool              `bson:"nivela4,omitempty"`
	ValvulaIn *bool              `bson:"valvulain,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
}

// "*" se utiliza para declarar punteros a tipos de dato, hacee los campos opcional como estructura inicial.
// omitempty se utiliza para indicar que un campo debe ser omitido si su valor es el valor cero del tipo de dato correspondiente, el campo no se vería en la tabla si no tiene valor.
// validate:"required" hace que sea requerido

type Range struct {
	From string `json:"from"`
	To   string `json:"to"`
}
