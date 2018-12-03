package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fullFile, _ := ioutil.ReadFile("input/input.txt")
	allrectangles := strings.Split(string(fullFile), "\n")
	totalCount := 0
	var coordinates map[string][]string
	var ids map[string]bool
	coordinates = make(map[string][]string)
	ids = make(map[string]bool)
	for _, v := range allrectangles {
		id := strings.Split(v, " ")[0]
		ids[id] = true
		leftPadding, _ := strconv.Atoi(strings.Split(strings.Split(v, " ")[2], ",")[0])
		topPadding, _ := strconv.Atoi(strings.Split(strings.Split(v, ": ")[0], ",")[1])
		width, _ := strconv.Atoi(strings.Split(strings.Split(v, ": ")[1], "x")[0])
		height, _ := strconv.Atoi(strings.Split(strings.Split(v, ": ")[1], "x")[1])
		for h := 1; h <= height; h++ {
			for w := 1; w <= width; w++ {
				coord := strconv.Itoa(leftPadding+w) + "x" + strconv.Itoa(topPadding+h) + "y"
				coordinates[coord] = append(coordinates[coord], id)
			}
		}
	}
	for _, v := range coordinates {
		if len(v) > 1 {
			totalCount++
			for _, id := range v {
				ids[id] = false
			}
		}
	}
	for k, v := range ids {
		if v {
			println(k)
		}
	}
	println(totalCount)
}
