package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
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
	series := trasformToIntArray(seriesRaw)

	return series
}

func isValid(prevSeries []int, number int) bool {
	for ai, a := range prevSeries {
		for _, b := range prevSeries[ai:] {
			if a+b == number {
				return true
			}
		}
	}

	return false
}

func findFirstInvalid(series []int, preamble int) int {
	for i, number := range series[preamble:] {
		validateSeries := series[i : i+preamble]
		if !isValid(validateSeries, number) {
			return number
		}
	}

	return -1
}

func arraySum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum = sum + i
	}
	return sum
}

func arrayMinMax(arr []int) (int, int) {
	min := math.MaxInt32
	max := math.MinInt32

	for _, i := range arr {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}
	return min, max
}

func findFirstSequenceCode(series []int, expectedSum int) int {
	for start := range series {
		offset := 1
		for true {
			if start+offset >= len(series) {
				break
			}
			sequence := series[start : start+offset]
			sequenceSum := arraySum(sequence)
			if sequenceSum == expectedSum {
				min, max := arrayMinMax(sequence)
				return min + max
			}

			offset = offset + 1
		}
	}

	return -1
}

func main() {
	series := prepareData("example.txt")
	firstInvalid := findFirstInvalid(series, 5)
	fmt.Println("first invalid number (example)", firstInvalid)
	sequenceCode := findFirstSequenceCode(series, firstInvalid)
	fmt.Println("first sequence code (example)", sequenceCode)

	series = prepareData("input.txt")
	firstInvalid = findFirstInvalid(series, 25)
	fmt.Println("first invalid number", firstInvalid)
	sequenceCode = findFirstSequenceCode(series, firstInvalid)
	fmt.Println("first sequence code", sequenceCode)
}
