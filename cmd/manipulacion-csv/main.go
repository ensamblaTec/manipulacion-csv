package main

import (
	"fmt"

	"github.com/g-code99/learning-path/week1/manipulacion-csv/pkg/utils"
)

func main() {
	var in string

	fmt.Printf("============ MENU ============\n" +
		"1. Listar ficheros\n2. Eliminar fichero\n3. Seleccionar fichero\n4. Salir\n>>> ")

	utils.ReadInput(&in)

	switch in {
	case "1":
		fmt.Println("Listar ficheros")
	case "2":
		fmt.Println("")
	case "3":
		fmt.Println("")
	case "4":
		fmt.Println("")
	default:
		fmt.Println("")
	}

	fmt.Println(in)
}
