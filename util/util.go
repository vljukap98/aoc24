package util

import (
	"strconv"
)

func ContainsI(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsS(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsR(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ArrStrToI(arr []string) []int {
	var arrInt []int

	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		arrInt = append(arrInt, j)
	}

	return arrInt
}

func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
