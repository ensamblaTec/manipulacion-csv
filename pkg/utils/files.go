package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func ReadFile(fileName string) {
	table, _ := tablewriter.NewCSV(os.Stdout, ("./data/" + fileName + ".csv"), true)
	table.Render()
}

// Leer los ficheros de un directorio
func ReadDir() {
	dir, err := os.ReadDir("./data")
	if os.IsExist(err) {
		log.Println("Doesn't exists directory data")
	}
	if err != nil {
		log.Println("Something has occured:", err)
	}
	for index, directory := range dir {
		fmt.Printf("%d. %s\n", (index + 1), directory.Name())
	}
}

func IsExists(fileName string) bool {
	_, err := os.Stat("./data/" + fileName + ".csv")
	if os.IsNotExist(err) {
		log.Println("File doesn't exists")
		return false
	}
	if err != nil {
		log.Println("Something has occured")
		return false
	}
	return true
}

func CreateFile(fileName string) {
	if len(fileName) == 0 {
		log.Println("Check the name")
	}
	file, err := os.Create("./data/" + fileName + ".csv")
	if err != nil {
		log.Println("Cannot create file")
	}
	defer file.Close()
	log.Println("The file:", file.Name(), "has been created.")
}

// Leer encabezados de archivos csv
func ReadHeaders(fileName string) {
	// primero abrir un archivo
	file, err := os.Open("./data/" + fileName + ".csv")
	if err != nil {
		log.Println("Cannot open file")
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.Read()

	if err != nil {
		log.Println("Cannot get content file")
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader(records)
	table.Render()
}

func WriteFile(fileName string, content [][]string) {
	filePath := "./data/" + fileName + ".csv" // Reemplaza esto con la ruta de tu archivo CSV

	// Abrir el archivo CSV en modo "append"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crear un escritor CSV
	writer := csv.NewWriter(file)

	// Agregar los nuevos datos al final del archivo CSV
	writer.WriteAll(content)

	// Escribir cualquier cambio pendiente al archivo
	writer.Flush()

	if err := writer.Error(); err != nil {
		fmt.Println("Error al escribir en el archivo CSV:", err)
		return
	}

	fmt.Println("Nuevos datos agregados exitosamente al archivo CSV.")
}
