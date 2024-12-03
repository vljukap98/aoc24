package util

import "strconv"

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
