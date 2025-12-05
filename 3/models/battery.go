package models

import "fmt"

type Battery struct {
	Banks []Bank
}

func (b *Battery) CalculateOutput() int {
	result := 0
	for i := range b.Banks {
		b.Banks[i].CalculateProduce()
		result = result + b.Banks[i].produce
	}
	return result
}

func (b *Battery) CalculateBigOutput() int {
	result := 0
	for i := range b.Banks {
		b.Banks[i].CalculateBigProduce()
		fmt.Println("Bank ", i, " has produce: ", b.Banks[i].produce)
		result = result + b.Banks[i].produce
	}
	return result
}

func CreateBattery() Battery {
	return Battery{}
}
