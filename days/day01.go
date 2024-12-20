package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func RunDay1Part1() {
	firstList, secondList := parseDay1()

	sort.Ints(firstList)
	sort.Ints(secondList)

	var diffs []int

	for i := 0; i < len(firstList); i++ {
		firstNum := firstList[i]
		secondNum := secondList[i]

		if firstNum > secondNum {
			diffs = append(diffs, firstNum-secondNum)
		} else {
			diffs = append(diffs, secondNum-firstNum)
		}
	}

	sum := 0
	for _, diff := range diffs {
		sum += diff
	}

	fmt.Println(sum)
}

func RunDay1Part2() {
	firstList, secondList := parseDay1()

	var scores []int

	for i := 0; i < len(firstList); i++ {
		firstNum := firstList[i]

		multiplier := 0
		for _, secondNum := range secondList {
			if secondNum == firstNum {
				multiplier++
			}
		}

		scores = append(scores, firstNum*multiplier)
	}

	sum := 0
	for _, score := range scores {
		sum += score
	}

	fmt.Println(sum)
}

func parseDay1() ([]int, []int) {
	// file, err := os.Open("inputs/day01_test.txt")
	file, err := os.Open("inputs/day01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var firstList []int
	var secondList []int

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		firstNum, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}
		firstList = append(firstList, firstNum)

		secondNum, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		secondList = append(secondList, secondNum)
	}
	return firstList, secondList
}
