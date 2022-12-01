package main

import (
	"aoc/data"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := data.GetData(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	arr := strings.Split(input, "\n")

	// part 1
	var maxCalories int
	var oneCalories int
	for _, calStr := range arr {
		if calStr == "" {
			if oneCalories > maxCalories {
				maxCalories = oneCalories
			}
			oneCalories = 0
			continue
		}

		cal, err := strconv.Atoi(calStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		oneCalories += cal
	}

	fmt.Println(maxCalories)

	// Part 2
	var topThereSum int
	calories := []int{}
	for _, calStr := range arr {
		if calStr == "" {
			if oneCalories > maxCalories {
				maxCalories = oneCalories
			}
			calories = append(calories, oneCalories)
			oneCalories = 0
			continue
		}

		cal, err := strconv.Atoi(calStr)
		if err != nil {
			fmt.Println(err)
			return
		}

		oneCalories += cal
	}

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})

	topThereSum = calories[0] + calories[1] + calories[2]
	fmt.Println(calories)
	fmt.Println(topThereSum)
}
