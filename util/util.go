package util

import (
	"fmt"
	"strconv"
)

func PrintMat(mat [][]string) {
	for _, v := range mat {
		fmt.Println(v)
	}
}

func Ator(text string) rune {
	var converter rune
	for _, runeValue := range text {
		converter = runeValue
	}
	return converter
}

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

func RemoveI(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func RemoveS(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
