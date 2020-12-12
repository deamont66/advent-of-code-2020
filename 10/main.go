package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkIndexArray(arr []string, index int) {
	if index < 0 || index >= len(arr) {
		err := errors.New("Error while parsing, trying to access undefined index")
		panic(err)
	}
}

func loadFileContent(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)

	return string(data)
}

func trasformToIntArray(numberSeriesRaw []string) []int {
	series := []int{}

	for _, numberRaw := range numberSeriesRaw {
		numberRaw = strings.TrimSpace(numberRaw)
		number, err := strconv.Atoi(numberRaw)
		check(err)

		series = append(series, number)
	}

	return series
}

func prepareData(filename string) []int {
	fileContent := loadFileContent(filename)
	fileContent = strings.TrimSpace(fileContent)

	seriesRaw := strings.Split(fileContent, "\n")
	adapters := trasformToIntArray(seriesRaw)

	adapters = append(adapters, 0) // starting with 0
	sort.Ints(adapters)

	adapters = append(adapters, adapters[len(adapters)-1]+3) // last is built-in +3 adapter

	return adapters
}

func calcAdapterTransformations(adapters []int) map[int]int {
	transformations := map[int]int{1: 0, 2: 0, 3: 0}

	for i, j := 0, 1; j < len(adapters); i, j = i+1, j+1 {
		ai := adapters[i]
		aj := adapters[j]
		transformations[aj-ai] = transformations[aj-ai] + 1
	}

	return transformations
}

func countAllPosibleConbinations(adapters []int, memo map[int]int) int {
	if len(adapters) == 1 {
		return 1
	}

	numberOfCombincations := 0
	fromAdapterValue := adapters[0]
	memoValue, found := memo[fromAdapterValue]
	if found {
		return memoValue
	}

	for i := 1; i < len(adapters); i++ {
		nextAdapterValue := adapters[i]
		if nextAdapterValue-fromAdapterValue > 3 {
			break
		}

		numberOfCombincations = numberOfCombincations + countAllPosibleConbinations(adapters[i:], memo)
	}

	memo[fromAdapterValue] = numberOfCombincations
	return numberOfCombincations
}

func run(filename string) {
	var memo = map[int]int{}
	adapters := prepareData(filename)
	adapterTransformations := calcAdapterTransformations(adapters)
	fmt.Println("transformations", filename, adapterTransformations)
	fmt.Println("result A", filename, adapterTransformations[1]*adapterTransformations[3])
	fmt.Println("cobinations to max", filename, countAllPosibleConbinations(adapters, memo))
}

func main() {
	run("example.txt")
	run("example1.txt")
	run("input.txt")
}
