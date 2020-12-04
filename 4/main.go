package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passportKey string

const (
	birthYear      passportKey = "byr"
	issueYear                  = "iyr"
	expirationYear             = "eyr"
	height                     = "hgt"
	hairColor                  = "hcl"
	eyeColor                   = "ecl"
	passportID                 = "pid"
	countryID                  = "cid"
)

var requiredKeys = [...]passportKey{
	birthYear,
	issueYear,
	expirationYear,
	height,
	hairColor,
	eyeColor,
	passportID,
	// countryID,
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

func createPassportMap(passportRaw string) map[passportKey]string {
	passport := map[passportKey]string{}

	keyValues := strings.Fields(passportRaw)

	for _, keyValue := range keyValues {
		split := strings.Split(keyValue, ":")
		checkIndexArray(split, 1)
		passport[passportKey(split[0])] = split[1]
	}

	return passport
}

func transformToPassportMap(passportsRaw []string) []map[passportKey]string {
	var passports []map[passportKey]string

	for _, str := range passportsRaw {
		if str == "" {
			continue
		}

		passport := createPassportMap(str)
		passports = append(passports, passport)
	}

	return passports
}

func prepareData() []map[passportKey]string {
	data := loadFileContent("input.txt")
	passportsRaw := strings.Split(data, "\n\n")

	return transformToPassportMap(passportsRaw)
}

func isPassportValid(passport map[passportKey]string) bool {
	for _, key := range requiredKeys {
		if _, ok := passport[key]; !ok {
			return false
		}
	}

	return true
}

func getValidPassportsCount(passports []map[passportKey]string) int {
	count := 0

	for _, passport := range passports {
		if isPassportValid(passport) {
			count = count + 1
		}
	}

	return count
}

func isPassportValidString(passport map[passportKey]string) bool {
	if !isPassportValid(passport) {
		return false
	}

	birthday, error := strconv.Atoi(passport[birthYear])
	check(error)
	if birthday < 1920 || birthday > 2002 {
		return false
	}

	issue, error := strconv.Atoi(passport[issueYear])
	check(error)
	if issue < 2010 || issue > 2020 {
		return false
	}

	expiration, error := strconv.Atoi(passport[expirationYear])
	check(error)
	if expiration < 2020 || expiration > 2030 {
		return false
	}

	heightReg, error := regexp.Compile(`^(\d+)(cm|in)$`)
	check(error)
	heightCorrect := heightReg.MatchString(passport[height])
	if !heightCorrect {
		return false
	}
	groups := heightReg.FindStringSubmatch(passport[height])
	if groups[2] == "cm" {
		size, error := strconv.Atoi(groups[1])
		check(error)

		if size < 150 || size > 193 {
			return false
		}
	} else if groups[2] == "in" {
		size, error := strconv.Atoi(groups[1])
		check(error)

		if size < 59 || size > 76 {
			return false
		}
	}

	hair, error := regexp.MatchString(`^#[0-9a-f]{6}$`, passport[hairColor])
	check(error)
	if !hair {
		return false
	}

	eye, error := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, passport[eyeColor])
	check(error)
	if !eye {
		return false
	}

	id, error := regexp.MatchString(`^\d{9}$`, passport[passportID])
	check(error)
	if !id {
		return false
	}

	return true
}

func getValidPassportsCountStrinct(passports []map[passportKey]string) int {
	count := 0

	for _, passport := range passports {
		if isPassportValidString(passport) {
			count = count + 1
		}
	}

	return count
}

func main() {
	passports := prepareData()

	fmt.Println("number of passports with all required fields", getValidPassportsCount(passports))
	fmt.Println("valid passports", getValidPassportsCountStrinct(passports))
}
