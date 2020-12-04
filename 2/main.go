package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	letter byte
	min    int
	max    int
}

type Password struct {
	value  string
	policy PasswordPolicy
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

func checkIndexString(arr string, index int) {
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

func createPassword(line string) Password {
	data := strings.Split(line, ":")
	checkIndexArray(data, 0)
	policyData := strings.Split(data[0], " ")
	checkIndexArray(policyData, 0)
	policyOccurenceData := strings.Split(policyData[0], "-")
	checkIndexArray(policyOccurenceData, 0)
	policyMin, err := strconv.Atoi(policyOccurenceData[0])
	check(err)
	checkIndexArray(policyOccurenceData, 1)
	policyMax, err := strconv.Atoi(policyOccurenceData[1])
	check(err)
	checkIndexArray(policyData, 1)
	checkIndexString(policyData[1], 0)
	policyLetter := policyData[1][0]
	checkIndexArray(data, 1)
	passwordValue := strings.TrimSpace(data[1])

	policy := PasswordPolicy{letter: policyLetter, min: policyMin, max: policyMax}
	password := Password{value: passwordValue, policy: policy}

	return password
}

func transformToPasswords(lines []string) []Password {
	var data []Password

	for _, str := range lines {
		if str == "" {
			continue
		}

		password := createPassword(str)
		data = append(data, password)
	}

	return data
}

func prepareData() []Password {
	data := loadFileContent("input.txt")
	lines := strings.Split(data, "\n")

	return transformToPasswords(lines)
}

func isValidOld(password Password) bool {
	count := strings.Count(password.value, string(password.policy.letter))

	return password.policy.min <= count && count <= password.policy.max
}

func countOldValidPasswords(passwords []Password) int {
	var count int = 0
	for _, password := range passwords {
		if isValidOld(password) {
			count = count + 1
		}
	}

	return count
}

func isValidNew(password Password) bool {
	isFirst := password.value[password.policy.min-1] == password.policy.letter
	isSecond := password.value[password.policy.max-1] == password.policy.letter

	return isFirst != isSecond
}

func countNewValidPasswords(passwords []Password) int {
	var count int = 0
	for _, password := range passwords {
		if isValidNew(password) {
			count = count + 1
		}
	}

	return count
}

func main() {
	passwords := prepareData()

	countOld := countOldValidPasswords(passwords)
	fmt.Println("valid passwords count (old)", countOld)

	countNew := countNewValidPasswords(passwords)
	fmt.Println("valid passwords count (new)", countNew)
}
