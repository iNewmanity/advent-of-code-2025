package models

import (
	"slices"
	"strconv"
)

type Bank struct {
	joltages []int
	produce  int
}

type ResultSet struct {
	value int
	index int
}

func CreateBank(sJoltages []string) Bank {
	var bank Bank
	for i := range sJoltages {
		joltage, _ := strconv.Atoi(sJoltages[i])
		bank.joltages = append(bank.joltages, joltage)
	}
	return bank
}

func (b *Bank) CalculateProduce() {
	var combinations []int
	for i := range b.joltages {
		sub := b.joltages[i+1:]
		for j := range sub {
			combinations = append(combinations, calculateProduce(b.joltages[i], sub[j]))
		}
	}
	b.produce = slices.Max(combinations)
}

func (b *Bank) CalculateBigProduce() {
	b.dissolveLowerBatteries()
	b.produce = b.calculateCompleteProduce()

}

func (b *Bank) calculateCompleteProduce() int {
	resultString := ""

	for _, value := range b.joltages {
		resultString += strconv.Itoa(value)
	}
	resultInt, _ := strconv.Atoi(resultString)
	return resultInt
}

func (b *Bank) dissolveLowerBatteries() {
	targetCount := 12
	removeCount := len(b.joltages) - targetCount

	stack := make([]int, 0, len(b.joltages))

	for _, value := range b.joltages {
		for removeCount > 0 && len(stack) > 0 && stack[len(stack)-1] < value {
			stack = stack[:len(stack)-1]
			removeCount--
		}
		stack = append(stack, value)
	}

	if len(stack) > targetCount {
		stack = stack[:targetCount]
	}

	b.joltages = stack
}

func (b *Bank) RemoveByIndex(index int) {
	b.joltages = append(b.joltages[:index], b.joltages[index+1:]...)
}

func (b *Bank) getMinWithIndex() (int, int) {
	result := slices.Min(b.joltages)
	indexofresult := slices.Index(b.joltages, result)
	return result, indexofresult
}

func calculateProduce(numberOne int, numberTwo int) int {
	result := (numberOne * 10) + numberTwo
	return result
}
