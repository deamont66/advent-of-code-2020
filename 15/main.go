package main

import (
	"errors"
	"fmt"
	"io/ioutil"
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

func parseBootstrapNumbers(numbersRaw []string) (map[int][]int, int) {
	history := map[int][]int{}
	lastNumber := 0

	for i, numberRaw := range numbersRaw {
		number, err := strconv.Atoi(numberRaw)
		check(err)

		lastNumber = number
		history[number] = append(history[number], i+1)
	}

	return history, lastNumber
}

func run(filename string, lastTurn int) {
	data := loadFileContent(filename)
	data = strings.TrimSpace(data)
	boostrapSequence := strings.Split(data, ",")

	turn := len(boostrapSequence)
	history, last := parseBootstrapNumbers(boostrapSequence)

	for turn < lastTurn {
		turn++
		next := 0
		if len(history[last]) != 1 {
			prev := 0
			if len(history[last]) > 1 {
				prev = history[last][len(history[last])-2]
			}
			next = history[last][len(history[last])-1] - prev
		}
		last = next
		history[next] = append(history[next], turn)
	}

	fmt.Println("number at turn ", lastTurn, filename, last)
}

func main() {
	run("example.txt", 2020)
	run("example.txt", 30000000)
	run("input.txt", 2020)
	run("input.txt", 30000000)
}
