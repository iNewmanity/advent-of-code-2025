package models

import (
	"strconv"
	"strings"
)

type Range struct {
	Start ID
	End   ID
}

func (r *Range) CreateIterableRange() IterableRange {
	ir := IterableRange{}
	for i := r.Start; i <= r.End; i++ {
		ir.iterableIDs = append(ir.iterableIDs, i)
	}
	return ir
}

func StringToRange(s string) Range {
	resultArray := strings.Split(s, "-")
	first, _ := strconv.Atoi(resultArray[0])
	last, _ := strconv.Atoi(resultArray[1])

	result := Range{Start: ToID(first), End: ToID(last)}
	return result

}
