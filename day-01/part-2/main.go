package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../actual.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	list_one := []int{}
	list_two := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		value_one := get_int(values[0])
		value_two := get_int(values[1])
		list_one = append(list_one, value_one)
		list_two = append(list_two, value_two)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	score_by_value := make(map[int]int)
	sum := 0

	for _, value_one := range list_one {
		score, exists := score_by_value[value_one]
		if exists {
			sum += score
		} else {
			number_found := get_occurences(value_one, list_two)
			score_by_value[value_one] = value_one * number_found
			sum += score_by_value[value_one]
		}
	}

	fmt.Println(sum)
}

func get_occurences(target int, int_arr []int) int {
	found := 0
	for _, value := range int_arr {
		if value == target {
			found++
		}
	}

	return found
}

func get_int(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		value = 0
	}

	return value
}

func absolute_value(i int) int {
	if i < 0 {
		i = -i
	}

	return i
}

func sum_array(int_arr []int) int {
	sum := 0
	for _, value := range int_arr {
		sum += value
	}

	return sum
}
