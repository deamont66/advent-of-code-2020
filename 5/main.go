package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type boardingPass struct {
	row    uint
	column uint
}

func (b boardingPass) seatID() uint {
	return b.row*8 + b.column
}

func newBoardingPass(passCode string) boardingPass {
	var row uint
	for _, c := range passCode[:7] {
		row = row << 1
		if c == 'B' {
			row = row | 01
		}
	}

	var column uint
	for _, c := range passCode[7:] {
		column = column << 1
		if c == 'R' {
			column = column | 01
		}
	}

	return boardingPass{row: row, column: column}
}

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

func transformToBoardingPasses(lines []string) []boardingPass {
	var passes []boardingPass

	for _, line := range lines {
		if line == "" {
			continue
		}
		passes = append(passes, newBoardingPass(line))
	}

	return passes
}

func prepareData() []boardingPass {
	data := loadFileContent("input.txt")
	lines := strings.Split(data, "\n")

	return transformToBoardingPasses(lines)
}

func getMaxBoardinPassID(passes []boardingPass) uint {
	var max uint
	for _, pass := range passes {
		currentID := pass.seatID()
		if max < currentID {
			max = currentID
		}
	}

	return max
}

func findMissingPassID(passes []boardingPass) (int, error) {
	var passIDs []int
	for _, pass := range passes {
		passIDs = append(passIDs, int(pass.seatID()))
	}
	sort.Ints(passIDs)

	numberOfPasses := len(passIDs)
	for i, id := range passIDs {
		if i < numberOfPasses-1 && passIDs[i+1] != id+1 {
			return id + 1, nil
		}
	}

	return 0, errors.New("Missing boarding pass not found")
}

func main() {
	boardingPasses := prepareData()

	fmt.Println("maximum pass ID", getMaxBoardinPassID(boardingPasses))

	missingPassID, err := findMissingPassID(boardingPasses)
	check(err)
	fmt.Println("missing pass ID", missingPassID)
}
