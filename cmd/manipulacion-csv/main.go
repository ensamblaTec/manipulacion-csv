package main

import (
	"fmt"
	"os"

	"github.com/g-code99/learning-path/week1/manipulacion-csv/pkg/utils"
)

func main() {
	var in string
	var inRead string
	var fileName string
	var inFile string

	for {
		fmt.Printf("============ MENU ============\n" +
			"1. Listar\n2. Crear\n3. Seleccionar\n4. Salir\n>>> ")
		utils.ReadInput(&in)
		switch in {
		case "1":
			fmt.Println("============ LISTANDO ============")
			utils.ReadDir()
			fmt.Println()
		case "2":
			fmt.Println("============ CREANDO ============")
			fmt.Print("Ingresa el nombre del archivo: ")
			utils.ReadInput(&inRead)
			utils.CreateFile(inRead)
			fmt.Println()
		case "3":
			fmt.Println("============ SELECCIONAR ============")
			fmt.Println("Escribe el nombre del archivo")
			fmt.Print(">>> ")
			utils.ReadInput(&fileName)
			if !utils.IsExists(fileName) {
				continue
			}
		for_select:
			for {
				fmt.Printf("============ " + fileName + " ============\n" +
					"1. Leer\n2. Eliminar\n3. Escribir\n4. Salir\n>>> ")
				utils.ReadInput(&inRead)
				switch inRead {
				case "1":
					fmt.Println("============ LEER ============")
					utils.ReadFile(fileName)
					fmt.Println()
				case "2":
					fmt.Println("============ ELIMINAR ============")
					fmt.Println("¿Estás seguro que deseas eliminar el archivo? s/n")
					utils.ReadInput(&inFile)
					if inFile == "s" {
						os.Remove("./data/" + fileName)
						break for_select
					}
					continue
				case "3":
					fmt.Println("============ ESCRIBIR ============")
					fmt.Println()
				case "4":
					fmt.Println("Salir...")
					fileName = ""
					inRead = ""
					break for_select
				default:
					fmt.Println("Opción incorrecta. Vuelva a intentar.")
				}
			}
			fmt.Println()
		case "4":
			fmt.Println("Saliendo...")
			os.Exit(1)
		default:
			fmt.Println("Opción incorrecta. Vuelva a intentar.")
			fmt.Println()
		}
	}
}
