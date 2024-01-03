package serialDataReception

import (
	"crs/models"
	"fmt"
)

func SerialDataReceptionActionHandler(in models.All, crs models.CrsInitConfigModel) {

	////////////////////////////////////////////////////////////////////
	// ACA COLOCAR TODA LA LOGICA DE COMPARACIÓN Y ALMACENADO EN BD
	////////////////////////////////////////////////////////////////////
	if in.B1Temp != nil {
		fmt.Println("*in.B1Temp", *in.B1Temp)
	} else {
		fmt.Println("in B1Temp is nil")
	}

	fmt.Println("in", in)
	fmt.Println("crs", crs)
}
