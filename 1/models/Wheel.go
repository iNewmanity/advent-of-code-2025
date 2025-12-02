package models

import "fmt"

type Wheel struct {
	Size               int
	CurrentPosition    int
	CountOfMagicNumber int
	MagicNumber        int
}

func (w *Wheel) ApplyRotations(r []Rotation) {
	for i := range r {
		fmt.Println("Rotating ", r[i], " steps")
		w.rotate(r[i])
	}
	fmt.Println("Current Position: ", w.CurrentPosition, " Count of magic number: ", w.CountOfMagicNumber, " Magic Number: ", w.MagicNumber)
}

func (w *Wheel) ApplyRotationsWatchingEveryMagicNumber(r []Rotation) {
	for i := range r {
		fmt.Println("Rotating ", r[i], " steps")
		if r[i].Value > w.Size {
			w.countStepsThatExceedSize(r[i])
			r[i].Cleanup(w)
			w.rotateSpecial(r[i])
		} else {
			w.rotateSpecial(r[i])
		}
	}
	fmt.Println("Current Position: ", w.CurrentPosition, " Count of magic number: ", w.CountOfMagicNumber, " Magic Number: ", w.MagicNumber)
}

func (w *Wheel) ResetPosition() {
	w.CurrentPosition = 50
}

func (w *Wheel) Reset() {
	w.CurrentPosition = 50
	w.CountOfMagicNumber = 0
}

func (w *Wheel) countStepsThatExceedSize(r Rotation) {
	w.CountOfMagicNumber = w.CountOfMagicNumber + r.RotationsWithoutCleanup(w)
	fmt.Println("...After Cleanup: ", r.Value, " Current Position: ", w.CurrentPosition, " Count of magic number: ", w.CountOfMagicNumber)
}

func (w *Wheel) rotateSpecial(r Rotation) {
	MAX_STEP := 99
	MIN_STEP := 0
	if r.Direction == "L" {
		if w.CurrentPosition-r.Value < MIN_STEP {
			w.CountOfMagicNumber++
		}
	} else if r.Direction == "R" {
		if w.CurrentPosition+r.Value > MAX_STEP {
			w.CountOfMagicNumber++
		}
	} else {
		panic("Invalid Direction")
	}
}

func (w *Wheel) rotate(r Rotation) {
	MAX_STEP := 99
	MIN_STEP := 0
	r.Cleanup(w)
	if r.Direction == "L" {
		w.CurrentPosition = w.CurrentPosition - r.Value
		if w.CurrentPosition == MIN_STEP {
			w.CountOfMagicNumber++
		} else if w.CurrentPosition < MIN_STEP {
			w.CurrentPosition = w.CurrentPosition + w.Size
			if w.CurrentPosition == MIN_STEP {
				w.CountOfMagicNumber++
			}
		}
	} else if r.Direction == "R" {
		w.CurrentPosition = w.CurrentPosition + r.Value
		if w.CurrentPosition == MIN_STEP {
			w.CountOfMagicNumber++
		} else if w.CurrentPosition > MAX_STEP {
			w.CurrentPosition = w.CurrentPosition - w.Size
			if w.CurrentPosition == MIN_STEP {
				w.CountOfMagicNumber++
			}
		}

	} else {
		panic("Invalid Direction")
	}

	fmt.Println("Number of Magic Numbers: ", w.CountOfMagicNumber, " Current Position: ", w.CurrentPosition, "")

}
