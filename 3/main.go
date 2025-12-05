package main

import (
	"advent-of-code-2025/3/models"
	"advent-of-code-2025/util"
	"fmt"
)

func main() {
	assignmentPath := "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2025/3/input/assignment.txt"
	data, _ := util.ReadFile(assignmentPath, "", true)

	var b models.Battery

	for i := range data {
		b.Banks = append(b.Banks, models.CreateBank(data[i]))
	}

	assignment1(b)

	assignment2(b)
}

func assignment1(b models.Battery) {
	fmt.Println(b.CalculateOutput())
}

func assignment2(b models.Battery) {
	fmt.Println(b.CalculateBigOutput())
}
