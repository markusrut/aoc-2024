package days

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/markusrut/aoc-2024/utils"
)

func RunDay2Part1() {
	reports := parseDay2()

	safeCount := 0
	for _, report := range reports {
		if isSafePart1(report) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func RunDay2Part2() {
	reports := parseDay2()

	safeCount := 0
	for _, report := range reports {
		if isSafePart2(report) {
			safeCount++
		}
	}

	fmt.Println(safeCount)
}

func isSafePart1(levels []int) bool {
	isIncreasing := levels[0] <= levels[1]

	isSafe := true
	for i := 0; i < len(levels)-1; i++ {
		if isIncreasing {
			diff := levels[i+1] - levels[i]
			if diff > 3 || diff <= 0 {
				isSafe = false
			}
		} else {

			diff := levels[i] - levels[i+1]
			if diff > 3 || diff <= 0 {
				isSafe = false
			}
		}
	}

	// fmt.Println(levels, isIncreasing, isSafe)

	return isSafe
}

func isSafePart2(levels []int) bool {
	isIncreasing := levels[0] <= levels[1]

	unsafeCount := 0
	for i := 0; i < len(levels)-1; i++ {
		if isIncreasing {
			diff := levels[i+1] - levels[i]
			if diff > 3 || diff <= 0 {
				unsafeCount++
			}
		} else {

			diff := levels[i] - levels[i+1]
			if diff > 3 || diff <= 0 {
				unsafeCount++
			}
		}
	}

	// fmt.Println(levels, isIncreasing, unsafeCount)

	return unsafeCount <= 1
}

func parseDay2() [][]int {
	// file, err := os.Open("inputs/day02_test.txt")
	file, err := os.Open("inputs/day02.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int

	for scanner.Scan() {
		report := strings.Fields(scanner.Text())

		intReport := utils.ConvertStringsToInts(report)

		reports = append(reports, intReport)
	}

	return reports
}
