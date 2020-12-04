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

func loadFileContent(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)

	return string(data)
}

func transformToInt(strings []string) []int {
	var data []int

	for _, str := range strings {
		if str == "" {
			continue
		}

		number, err := strconv.Atoi(str)
		check(err)
		data = append(data, number)
	}

	return data
}

func prepareData() []int {
	data := loadFileContent("input.txt")
	lines := strings.Split(data, "\n")
	numbers := transformToInt(lines)

	return numbers
}

func getTwoExpenses(numbers []int) (int, error) {
	for i, num1 := range numbers {
		for j := i; j < len(numbers); j++ {
			num2 := numbers[j]
			if num1+num2 == 2020 {
				return num1 * num2, nil
			}
		}
	}

	return 0, errors.New("could not found two numbers in array equal to 2020")
}

func getThreeExpenses(numbers []int) (int, error) {
	for i, num1 := range numbers {
		for j := i; j < len(numbers); j++ {
			num2 := numbers[j]
			for k := j; k < len(numbers); k++ {
				num3 := numbers[k]
				if num1+num2+num3 == 2020 {
					return num1 * num2 * num3, nil
				}
			}
		}
	}

	return 0, errors.New("could not found three numbers in array equal to 2020")
}

func getTwoExpensesRec(numbers []int) (int, error) {
	return rec2(numbers, 0, 1)
}

func rec2(numbers []int, i int, j int) (int, error) {
	if i >= len(numbers) {
		return 0, errors.New("could not found two numbers in array equal to 2020")
	}

	if numbers[i]+numbers[j] == 2020 {
		return numbers[i] * numbers[j], nil
	}

	if j == len(numbers)-1 {
		newI := i + 1
		return rec2(numbers, newI, newI)
	}

	return rec2(numbers, i, j+1)
}

func getThreeExpensesRec(numbers []int) (int, error) {
	return rec3(numbers, 0, 1, 2)
}

func rec3(numbers []int, i int, j int, k int) (int, error) {
	if i >= len(numbers) {
		return 0, errors.New("could not found three numbers in array equal to 2020")
	}

	if j >= len(numbers) {
		newI := i + 1
		return rec3(numbers, newI, newI+1, newI+2)
	}

	if k >= len(numbers) {
		newJ := j + 1
		return rec3(numbers, i, newJ, newJ+1)
	}

	if numbers[i]+numbers[j]+numbers[k] == 2020 {
		return numbers[i] * numbers[j] * numbers[k], nil
	}

	return rec3(numbers, i, j, k+1)
}

func main() {
	numbers := prepareData()

	twoExpensesMultiple, err := getTwoExpenses(numbers)
	check(err)
	fmt.Println("two", twoExpensesMultiple)

	threeExpensesMultiple, err := getThreeExpenses(numbers)
	check(err)
	fmt.Println("three", threeExpensesMultiple)

	twoExpensesMultipleRec, err := getTwoExpensesRec(numbers)
	check(err)
	fmt.Println("two_rec", twoExpensesMultipleRec)

	threeExpensesMultipleRec, err := getThreeExpensesRec(numbers)
	check(err)
	fmt.Println("three_rec", threeExpensesMultipleRec)
}
