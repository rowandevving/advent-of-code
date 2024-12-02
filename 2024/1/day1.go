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
	input := goaocd.Input(2024, 1)

	list1, list2 := parse(input)

	answer1 := part1(list1, list2)
	answer2 := part2(list1, list2)

	fmt.Printf("Part 1: %d\nPart 2: %d", answer1, answer2)
}

func parse(input string) ([]int, []int) {
	list1, list2 := []int{}, []int{}

	for _, line := range strings.Split(input, "\n") {

		if line == "" {
			continue
		}

		nums := strings.Fields(line)

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}

func part1(list1 []int, list2 []int) int {

	slices.Sort(list1)
	slices.Sort(list2)

	diff := 0
	for i := 0; i < len(list1); i++ {
		diff += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return diff
}

func part2(list1 []int, list2 []int) int {

	counts := make(map[int]int)
	for _, num := range list2 {
		counts[num]++
	}

	similarityScore := 0

	for _, num := range list1 {
		similarityScore += num * counts[num]
	}

	return similarityScore
}
