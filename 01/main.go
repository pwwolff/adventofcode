package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var frequency int = 0
	var changeOfFrequency int
	var frequencies map[int]int
	doubleFrequencyFound := false
	directions := []int{}
	frequencies = make(map[int]int)
	frequencies[0] = 1
	file, err := os.Open("input/input.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := strings.Replace(scanner.Text(), "+", "", 1)

		changeOfFrequency, _ = strconv.Atoi(currentLine)
		directions = append(directions, changeOfFrequency)
		frequency += changeOfFrequency
		frequencies[frequency] = 1
	}
	println("Star1: ", frequency)
	for !doubleFrequencyFound {
		for _, element := range directions {
			frequency += element
			if _, ok := frequencies[frequency]; ok {
				if !doubleFrequencyFound {
					println("Star2", frequency)
				}
				doubleFrequencyFound = true
			} else {
				frequencies[frequency] = 1
			}
		}
	}

}
