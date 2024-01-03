package models

type CrsInitConfigModel struct {
	B1 struct {
		Termico     bool `json:"Termico"`
		Encendido   bool `json:"Encendido"`
		Temperatura struct {
			Check    bool    `json:"Check"`
			Cambio   bool    `json:"Cambio"`
			TimeLoop float64 `json:"TimeLoop"`
			Maxima   float64 `json:"Maxima"`
		} `json:"Temperatura"`
	} `json:"B1"`
	B2 struct {
		Termico     bool `json:"Termico"`
		Encendido   bool `json:"Encendido"`
		Temperatura struct {
			Check    bool    `json:"Check"`
			Cambio   bool    `json:"Cambio"`
			TimeLoop float64 `json:"TimeLoop"`
			Maxima   float64 `json:"Maxima"`
		} `json:"Temperatura"`
	} `json:"B2"`
	B3 struct {
		Termico     bool `json:"Termico"`
		Encendido   bool `json:"Encendido"`
		Temperatura struct {
			Check    bool    `json:"Check"`
			Cambio   bool    `json:"Cambio"`
			TimeLoop float64 `json:"TimeLoop"`
			Maxima   float64 `json:"Maxima"`
		} `json:"Temperatura"`
	} `json:"B3"`
	Altura struct {
		Check      bool    `json:"Check"`
		TimeLoop   float64 `json:"TimeLoop"`
		CadaCambio bool    `json:"CadaCambio"`
		H1         float64 `json:"H1"`
		H2         float64 `json:"H2"`
		H3         float64 `json:"H3"`
		H4         float64 `json:"H4"`
		H5         float64 `json:"H5"`
	} `json:"Altura"`
	Presion struct {
		Check             bool    `json:"Check"`
		CheckAntesBomba   bool    `json:"CheckAntesBomba"`
		CheckDespuesBomba bool    `json:"CheckDespuesBomba"`
		Min               float64 `json:"Min"`
		Max               float64 `json:"Max"`
		TimeLoop          float64 `json:"TimeLoop"`
	} `json:"Presion"`
	Cambios struct {
		P1 bool `json:"P1"`
		P2 bool `json:"P2"`
		P3 bool `json:"P3"`
		P4 bool `json:"P4"`
	} `json:"Cambios"`
	Fuente struct {
		Check      bool `json:"Check"`
		CadaCambio bool `json:"CadaCambio"`
	} `json:"Fuente"`
	Flujo struct {
		Calle struct {
			Check    bool    `json:"Check"`
			TimeLoop float64 `json:"TimeLoop"`
		} `json:"Calle"`
		Tanque struct {
			Check    bool    `json:"Check"`
			TimeLoop float64 `json:"TimeLoop"`
		} `json:"Tanque"`
	} `json:"Flujo"`
	Valvula struct {
		Check      bool `json:"Check"`
		CadaCambio bool `json:"CadaCambio"`
	} `json:"Valvula"`
	Comunicacion struct {
		Baud int    `json:"Baud"`
		Port string `json:"Port"`
	} `json:"Comunicacion"`
}
