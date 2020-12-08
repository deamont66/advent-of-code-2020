package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type operation struct {
	operationType     string
	operationArgument int
	visitedFlag       bool
}

func (current operation) run(op int, acc int) (int, int, error) {
	switch current.operationType {
	case "nop":
		return op + 1, acc, nil
	case "acc":
		return op + 1, acc + current.operationArgument, nil
	case "jmp":
		return op + current.operationArgument, acc, nil
	}

	return op, acc, errors.New("Invalid operation type " + current.operationType)
}

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

func newOperation(operationString string) operation {
	operationString = strings.TrimSpace(operationString)
	parts := strings.Split(operationString, " ")
	checkIndexArray(parts, 1)

	operationType := parts[0]
	operationArgument, err := strconv.Atoi(parts[1])
	check(err)

	return operation{operationType: operationType, operationArgument: operationArgument, visitedFlag: false}
}

func trasformToOperationArray(operationLines []string) []operation {
	operations := []operation{}

	for _, operationLine := range operationLines {
		operations = append(operations, newOperation(operationLine))
	}

	return operations
}

func prepareData(filename string) []operation {
	fileContent := loadFileContent(filename)
	fileContent = strings.TrimSpace(fileContent)

	operationsRaw := strings.Split(fileContent, "\n")
	operations := trasformToOperationArray(operationsRaw)

	return operations
}

func process(operations []operation, startOp int, startAcc int) (int, bool) {
	op := startOp
	acc := startAcc

	for op < len(operations) && !operations[op].visitedFlag {
		operations[op].visitedFlag = true
		newOp, newAcc, err := operations[op].run(op, acc)
		check(err)

		op = newOp
		acc = newAcc
	}

	return acc, op >= len(operations)
}

func clearOperationsVisited(operations []operation) []operation {
	for i := 0; i < len(operations); i++ {
		operations[i].visitedFlag = false
	}

	return operations
}

func try(operations []operation) (int, error) {
	for i := 0; i < len(operations); i++ {
		operations = clearOperationsVisited(operations)
		originalType := operations[i].operationType
		if originalType == "jmp" {
			operations[i].operationType = "nop"
		} else if originalType == "nop" {
			operations[i].operationType = "jmp"
		}
		acc, finished := process(operations, 0, 0)
		operations[i].operationType = originalType

		if finished {
			return acc, nil
		}
	}

	return 0, errors.New("Unable to finish")
}

func main() {
	operations := prepareData("example.txt")
	acc, _ := process(operations, 0, 0)
	fmt.Println("last acc value before infinite loop (example)", acc)
	finishedAcc, err := try(operations)
	check(err)
	fmt.Println("acc value after program finished (example)", finishedAcc)

	operations = prepareData("input.txt")
	acc, _ = process(operations, 0, 0)
	fmt.Println("last acc value before infinite loop", acc)
	finishedAcc, err = try(operations)
	check(err)
	fmt.Println("acc value after program finished", finishedAcc)
}
