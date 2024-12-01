package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	distances := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		value_one := get_int(values[0])
		value_two := get_int(values[1])
		list_one = append(list_one, value_one)
		list_two = append(list_two, value_two)
	}

	sort.Ints(list_one)
	sort.Ints(list_two)

	for index, value_one := range list_one {
		value_two := list_two[index]
		distance := value_one - value_two
		distances = append(distances, absolute_value(distance))
	}

	distance_sum := sum_array(distances)
	fmt.Println(distance_sum)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
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
