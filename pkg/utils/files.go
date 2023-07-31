package utils

import (
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
