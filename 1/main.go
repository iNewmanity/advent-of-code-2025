package main

import (
	"advent-of-code-2025/1/models"
	"advent-of-code-2025/util"
)

func main() {
	assignment := "/home/janneumann/Dokumente/Daten/Projekte/Privat/advent-of-code-2025/1/input/assignment_1.txt"
	_, data := util.ReadFile(assignment, "", false)

	wheel := models.Wheel{
		Size:               100,
		CurrentPosition:    50,
		CountOfMagicNumber: 0,
		MagicNumber:        0,
	}

	var rotations []models.Rotation

	for i := range data {
		rotations = append(rotations, models.ConvertStringToRotation(data[i]))
	}

	assignment1(rotations, wheel)
	assignment2(rotations, wheel)

}

func assignment1(rotations []models.Rotation, wheel models.Wheel) {
	wheel.ApplyRotations(rotations)
}

func assignment2(rotations []models.Rotation, wheel models.Wheel) {
	wheel.ApplyRotations(rotations)
	wheel.ResetPosition()
	wheel.ApplyRotationsWatchingEveryMagicNumber(rotations)
}
