package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/Olegas/goaocd"
)

func main() {
	input := goaocd.Input(2024, 2)

	reports := parse(input)

	answer1 := part1(reports)

	fmt.Printf("Part 1: %d\n", answer1)

}

func parse(input string) [][]int {
	reports := [][]int{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		nums := strings.Fields(line)
		report := []int{}

		for _, num := range nums {
			n, _ := strconv.Atoi(num)
			report = append(report, n)
		}

		reports = append(reports, report)
	}
	return reports
}

func part1(reports [][]int) int {

	safeReports := 0

	for _, report := range reports {
		if isDescending(report) || slices.IsSorted(report) {

			safe := true

			for i := 1; i < len(report); i++ {
				diff := math.Abs(float64(report[i] - report[i-1]))
				if !(1 <= diff && diff <= 3) {
					safe = false
					break
				}
			}

			if safe {
				safeReports++
			}
		}
	}

	return safeReports
}

func isDescending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			return false
		}
	}
	return true
}
