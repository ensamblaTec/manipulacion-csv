package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func ReadFile(fileName string) {
	table, _ := tablewriter.NewCSV(os.Stdout, ("testdata/" + fileName + ".csv"), true)

	table.Render()
}
