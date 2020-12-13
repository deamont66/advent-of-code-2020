package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	floor    = 0
	empty    = 1
	ocuppied = 2
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

func parseMapRow(mapRowRaw string) []int {
	row := []int{}
	for _, tile := range mapRowRaw {
		tileValue := floor
		if tile == 'L' {
			tileValue = ocuppied // first we need to occupie every seat
		} else if tile == '#' {
			tileValue = ocuppied
		}
		row = append(row, tileValue)
	}

	return row
}

func trasformToPlanMap(planRaw []string) [][]int {
	planMap := [][]int{}

	for _, planRowRaw := range planRaw {
		row := parseMapRow(planRowRaw)
		planMap = append(planMap, row)
	}

	return planMap
}

func prepareData(filename string) [][]int {
	fileContent := loadFileContent(filename)
	fileContent = strings.TrimSpace(fileContent)

	mapRaw := strings.Split(fileContent, "\n")
	planMap := trasformToPlanMap(mapRaw)

	return planMap
}

var nw = []int{-1, -1}
var n = []int{-1, 0}
var ne = []int{-1, 1}
var e = []int{0, 1}
var se = []int{1, 1}
var s = []int{1, 0}
var sw = []int{1, -1}
var w = []int{0, -1}

var sides = [][]int{nw, n, ne, e, se, s, sw, w}

func getNumberOfAdjecentOccupiedSeats(mapPlan [][]int, i int, j int, size int) int {
	numberOfOccupiedSeats := 0

	iMax := len(mapPlan)
	jMax := len(mapPlan[i])

	for _, side := range sides {
		iTest := i
		jTest := j
		iteraation := 0
		for iTest >= 0 && iTest < iMax && jTest >= 0 && jTest < jMax {
			if iteraation != 0 {
				if mapPlan[iTest][jTest] == ocuppied {
					numberOfOccupiedSeats++
					break
				} else if mapPlan[iTest][jTest] == empty {
					break
				}

				if size != 0 && iteraation >= size {
					break
				}
			}

			iTest = iTest + side[0]
			jTest = jTest + side[1]
			iteraation++
		}
	}

	return numberOfOccupiedSeats
}

func createNewMap(mapPlan [][]int) [][]int {
	planMap := [][]int{}
	for i := 0; i < len(mapPlan); i++ {
		newRow := []int{}
		for j := 0; j < len(mapPlan[i]); j++ {
			value := mapPlan[i][j]
			newRow = append(newRow, value)
		}
		planMap = append(planMap, newRow)
	}

	return planMap
}

func stepTheGame(mapPlan [][]int, checkSize int, leaveThreashold int) ([][]int, bool) {
	newMap := createNewMap(mapPlan)

	hasChanges := false
	for i := 0; i < len(mapPlan); i++ {
		for j := 0; j < len(mapPlan[i]); j++ {
			adjecentOccupiedSeats := getNumberOfAdjecentOccupiedSeats(mapPlan, i, j, checkSize)

			if mapPlan[i][j] == empty && adjecentOccupiedSeats == 0 {
				newMap[i][j] = ocuppied
				hasChanges = true
			} else if mapPlan[i][j] == ocuppied && adjecentOccupiedSeats >= leaveThreashold {
				newMap[i][j] = empty
				hasChanges = true
			}
		}
	}

	return newMap, hasChanges
}

func countOccupiedSeats(mapPlan [][]int) int {
	count := 0
	for i := 0; i < len(mapPlan); i++ {
		for j := 0; j < len(mapPlan[i]); j++ {
			if mapPlan[i][j] == ocuppied {
				count++
			}
		}
	}

	return count
}

func run(filename string, checkSize int, leaveThreashold int) {
	planMap := prepareData(filename)

	nextPlanMap, hasChanges := stepTheGame(planMap, checkSize, leaveThreashold)
	for hasChanges {
		nextPlanMap, hasChanges = stepTheGame(nextPlanMap, checkSize, leaveThreashold)
	}
	configuration := "(" + strconv.Itoa(checkSize) + ", " + strconv.Itoa(leaveThreashold) + ")"
	fmt.Println("number of occupied seats", filename, configuration, countOccupiedSeats(nextPlanMap))
}

func main() {
	run("example.txt", 1, 4)
	run("example.txt", 0, 5)
	run("input.txt", 1, 4)
	run("input.txt", 0, 5)
}
