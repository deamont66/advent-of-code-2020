package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func loadFileContent(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)

	return string(data)
}

func createMapLine(line string) []bool {
	var mapLine []bool
	for _, c := range line {
		mapLine = append(mapLine, c == '#')
	}

	return mapLine
}

func transformToMap(lines []string) [][]bool {
	var data [][]bool

	for _, str := range lines {
		if str == "" {
			continue
		}

		mapLine := createMapLine(str)
		data = append(data, mapLine)
	}

	return data
}

func prepareData() [][]bool {
	data := loadFileContent("input.txt")
	lines := strings.Split(data, "\n")

	return transformToMap(lines)
}

func move(mapData [][]bool, stepsRight int, stepsDown int, i int, j int, treesEncountred int) int {
	if i >= len(mapData) {
		return treesEncountred
	}

	if mapData[i][j] {
		treesEncountred = treesEncountred + 1
	}

	return move(mapData, stepsRight, stepsDown, i+stepsDown, (j+stepsRight)%len(mapData[i]), treesEncountred)
}

func main() {
	mapData := prepareData()

	trees11 := move(mapData, 1, 1, 0, 0, 0)
	fmt.Println("1-1 steps", trees11)

	trees31 := move(mapData, 3, 1, 0, 0, 0)
	fmt.Println("3-1 steps", trees31)

	trees51 := move(mapData, 5, 1, 0, 0, 0)
	fmt.Println("5-1 steps", trees51)

	trees71 := move(mapData, 7, 1, 0, 0, 0)
	fmt.Println("7-1 steps", trees71)

	trees12 := move(mapData, 1, 2, 0, 0, 0)
	fmt.Println("1-2 steps", trees12)

	fmt.Println("all slopes multiply", trees11*trees31*trees51*trees71*trees12)
}
