package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

type step struct {
	letter       string
	predecessors []string
	dependencies []string
}

func (s step) prettyPrint() {
	fmt.Printf("%s after: %v, before: %v\n", s.letter, s.dependencies, s.predecessors)
}

func main() {
	allSteps := ExtractData()
	doNext := GetPossible(allSteps)[0]
	done := DoStep(allSteps, doNext, "")
	println(done)
}

func DoStep(steps map[string]*step, toDo string, allDone string) string {
	delete(steps, toDo)
	allDone += toDo
	for _, v := range steps {
		for i, s := range v.predecessors {
			if s == toDo {
				v.predecessors = append(v.predecessors[:i], v.predecessors[i+1:]...)
			}
		}
	}
	doNext := GetPossible(steps)
	if len(doNext) == 0 {
		return allDone
	}
	return DoStep(steps, doNext[0], allDone)
}

func GetPossible(steps map[string]*step) []string {
	var possible []string
	for k, v := range steps {
		if len(v.predecessors) == 0 {
			possible = append(possible, k)
		}
	}
	sort.SliceStable(possible, func(i, j int) bool { return possible[i] < possible[j] })
	return possible
}

func ExtractData() map[string]*step {
	contents, _ := ioutil.ReadFile("input/input.txt")
	allInstructions := strings.Split(string(contents), "\n")
	var allSteps map[string]*step
	allSteps = make(map[string]*step)
	for _, val := range allInstructions {
		re, _ := regexp.Compile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin.`)
		results := re.FindAllStringSubmatch(val, -1)
		for _, v := range results {
			step1 := v[1]
			step2 := v[2]
			if _, ok := allSteps[step1]; !ok {
				allSteps[step1] = &step{letter: step1}
			}
			if _, ok := allSteps[step2]; !ok {
				allSteps[step2] = &step{letter: step2}
			}
			allSteps[step1].dependencies = append(allSteps[step1].dependencies, step2)
			allSteps[step2].predecessors = append(allSteps[step2].predecessors, step1)

		}
	}
	return (allSteps)
}
