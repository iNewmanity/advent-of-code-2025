package models

import (
	"strconv"
	"strings"
)

type ID int

func (i *ID) isInvalid() bool {
	iString := i.toString()
	if isEven(iString) {
		if isInvalidBy(2, iString) {
			return true
		} else if isInvalidBy(len(iString)/2, iString) {
			return true
		} else if isInvalidBySpecial(iString) {
			return true
		} else {
			return false
		}
	} else {
		if isInvalidBy(len(iString)/1, iString) {
			return true
		} else if len(iString)%3 == 0 && isInvalidBy(3, iString) {
			return true
		} else if len(iString)%5 == 0 && isInvalidBy(5, iString) {
			return true
		} else if isInvalidBySpecial(iString) {
			return true
		} else {
			return false
		}
	}
}

func isInvalidBy(i int, s string) bool {
	if i == 1 {
		if countOf(s, s[0:len(s)/i]) >= i+1 {
			return true
		} else {
			return false
		}
	} else {
		if countOf(s, s[0:len(s)/i]) == i {
			return true
		} else {
			return false
		}
	}

}

func isInvalidBySpecial(s string) bool {
	if countOf(s, firstChar(s)) == len(s) && len(s) > 1 {
		return true
	} else {
		return false
	}
}

func countOf(s string, c string) int {
	return strings.Count(s, c)
}

func firstChar(s string) string {
	return string(s[0])
}

func firstOneThird(s string) string {
	return s[:len(s)/3]
}

func extractFirstHalfString(s string) string {
	return s[:len(s)/2]
}

func isEven(s string) bool {
	return len(s)%2 == 0
}

func (i *ID) toString() string {
	return strconv.Itoa(int(*i))
}

func ToID(i int) ID {
	return ID(i)
}
