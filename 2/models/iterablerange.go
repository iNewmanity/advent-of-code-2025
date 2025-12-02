package models

import "fmt"

type IterableRange struct {
	iterableIDs []ID
	invalidIDs  []ID
}

func (ir *IterableRange) Iterate() {
	for i := range ir.iterableIDs {
		if ir.iterableIDs[i].isInvalid() == true {
			ir.invalidIDs = append(ir.invalidIDs, ir.iterableIDs[i])
		} else {

		}
	}
}

func (ir *IterableRange) PrintInvalidIDs() {
	fmt.Println(ir.invalidIDs)
}

func (ir *IterableRange) GetSumOfInvalidIDs() int {
	result := 0
	for i := range ir.invalidIDs {
		result = result + int(ir.invalidIDs[i])
	}
	return result
}
