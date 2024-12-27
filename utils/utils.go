package utils

import (
	"log"
	"strconv"
)

func ConvertStringsToInts(stringsList []string) []int {
	var intsList []int
	for _, str := range stringsList {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		intsList = append(intsList, num)
	}
	return intsList
}
