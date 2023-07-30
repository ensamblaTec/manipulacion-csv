package models

import (
	"errors"
	"fmt"
)

type CSVFile struct {
	Encabezado          []string
	NumFilas, NumColumn uint
	Filas               []string
}

func (file *CSVFile) SetEncabezado(encabezados []string) {
	file.Encabezado = encabezados
	file.NumColumn = uint(len(encabezados))
}

func (file *CSVFile) RemoveEncabezado() {
	file.Encabezado = []string{}
}

func (file *CSVFile) SetFilas(filas []string) {
	file.Filas = filas
	file.NumFilas = uint(len(filas))
}

func (file *CSVFile) RemoveFilas() {
	file.Filas = []string{}
}

func (file *CSVFile) AgregarColumna(encabezados []string) {
	file.Encabezado = append(file.Encabezado, encabezados...)
}

func (file *CSVFile) AgregarFila(filas []string) {
	file.Filas = append(file.Filas, filas...)
}

func (file *CSVFile) RemoverFila(fila int) error {
	if fila > len(file.Filas) || fila < 0 {
		return errors.New(fmt.Sprint("Cannot remove file ", fila))
	}
	file.Filas = append(file.Filas[:fila], file.Filas[:fila+1]...)
	return nil
}

func (file *CSVFile) Imprimir() {
	fmt.Println(file.Encabezado)
	for v := range file.Filas {
		fmt.Println(v)
	}
}
