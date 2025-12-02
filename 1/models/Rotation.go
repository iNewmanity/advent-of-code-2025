package models

import (
	"strconv"
	"strings"
)

type Rotation struct {
	Direction Direction
	Value     int
}

func (r *Rotation) Cleanup(w *Wheel) {
	r.Value = r.Value % w.Size
}

func (r *Rotation) RotationsWithoutCleanup(w *Wheel) int {
	rest := r.Value % w.Size
	return (r.Value - rest) / w.Size
}

func ConvertStringToRotation(input string) Rotation {
	return Rotation{
		Value:     GetSteps(input),
		Direction: Direction(input[0]),
	}
}

func GetSteps(input string) int {
	if strings.ContainsAny(input, "R") {
		result, found := strings.CutPrefix(input, "R")
		if found {
			result, _ := strconv.Atoi(result)
			return result
		}
	} else if strings.ContainsAny(input, "L") {
		result, found := strings.CutPrefix(input, "L")
		if found {
			result, _ := strconv.Atoi(result)
			return result
		}
	} else {
		panic("Invalid input")
	}
	return 0
}
