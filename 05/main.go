package main

import (
	"io/ioutil"
	"strings"
)

func main() {

	for i := 65; i < 91; i++ {
		CollapsePolymer("input/input.txt", i)
	}
}

func CollapsePolymer(fp string, removestring int) {
	full, _ := ioutil.ReadFile(fp)
	fullFile := strings.Replace(strings.Replace(string(full), string(removestring), "", -1), string(removestring+32), "", -1)
	previous := 0
	current := 0
	diff := 0

	for i := 1; i < len(fullFile); i++ {
		previous = int(fullFile[i-1])
		current = int(fullFile[i])
		diff = previous - current
		if diff == 32 || diff == -32 {
			fullFile = fullFile[0:i-1] + fullFile[i+1:]
			i = 1
		}
	}
	for i := 1; i < len(fullFile); i++ {
		previous = int(fullFile[i-1])
		current = int(fullFile[i])
		diff = previous - current
		if diff == 32 || diff == -32 {
			fullFile = fullFile[0:i-1] + fullFile[i+1:]
			i = 1
		}
	}
	println(string(removestring), len(fullFile))
}
