package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const minimum_difference int = 1
const maximum_difference int = 3

func main() {
	file, err := os.Open("../actual.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reports := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		report := get_ints(values)
		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	unsafe_count := 0

	for _, report := range reports {
		direction := 0
		previous_level := report[0]
		for index := 1; index < len(report); index++ {
			current_level := report[index]
			difference := current_level - previous_level
			if is_difference_in_bounds(difference) {
				if direction == 0 {
					direction = difference / absolute_value(difference)
				} else if direction != difference/absolute_value(difference) {
					unsafe_count++
					break
				}
			} else {
				unsafe_count++
				break
			}
			previous_level = current_level
		}
	}

	fmt.Println(len(reports) - unsafe_count)
}

func is_difference_in_bounds(level_difference int) bool {
	difference := absolute_value(level_difference)
	return difference >= minimum_difference && difference <= maximum_difference
}

func get_int(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		value = 0
	}

	return value
}

func get_ints(strs []string) []int {
	ints := []int{}
	for _, value := range strs {
		integer := get_int(value)
		ints = append(ints, integer)
	}

	return ints
}

func absolute_value(i int) int {
	if i < 0 {
		i = -i
	}

	return i
}
