package models_test

import (
	"testing"

	"github.com/g-code99/learning-path/week1/manipulacion-csv/internal/models"
)

func TestAgregarEncabezado(t *testing.T) {
	archivo := models.CSVFile{}

	archivo.AgregarEncabezado([]string{"Hola", "Mundo"})

	if !CompareSlice(archivo.Encabezado, []string{"Hola", "Mundo"}) {
		t.Fail()
	}
}
func TestRemoverEncabezado(t *testing.T) {
	archivo := models.CSVFile{}

	archivo.AgregarEncabezado([]string{"Hola", "Mundo"})

	archivo.RemoveEncabezado()

	if len(archivo.Encabezado) > 0 {
		t.Fail()
	}
}
func TestAgregarFila(t *testing.T) {
	archivo := models.CSVFile{}

	archivo.AgregarFila([]string{"Fila 1", "Dato 1"})
	archivo.AgregarFila([]string{"Fila 2", "Dato 2"})

	if !CompareFilas(archivo.Filas, [][]string{
		{"Fila 1", "Dato 1"},
		{"Fila 2", "Dato 2"},
	}) {
		t.Fail()
	}
}
func TestRemoverFila(t *testing.T) {
	archivo := models.CSVFile{}

	archivo.AgregarFila([]string{"Fila 1", "Dato 1"})
	archivo.AgregarFila([]string{"Fila 2", "Dato 2"})

	archivo.RemoveFilas()

	if len(archivo.Filas) > 0 {
		t.Fail()
	}
}

func CompareSlice(encabezado, testValue []string) bool {
	for i, v := range encabezado {
		if v != testValue[i] {
			return false
		}
	}
	return true
}

func CompareFilas(filas, testValue [][]string) bool {
	for i, v := range filas {
		for j, k := range v {
			if k != testValue[i][j] {
				return false
			}
		}
	}
	return true
}
