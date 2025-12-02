package main

import (
	"advent-of-code-2025/2/models"
	"advent-of-code-2025/util"
	"fmt"
)

func main() {
	assignmentpath := "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2025/2/input/assignment_1.txt"
	data, _ := util.ReadFile(assignmentpath, ",", true)
	realData := data[0]
	result := 0

	for i := range realData {
		r := models.StringToRange(realData[i])
		ir := r.CreateIterableRange()
		ir.Iterate()
		ir.PrintInvalidIDs()
		result = result + ir.GetSumOfInvalidIDs()
	}

	fmt.Println(result)
}
