package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type PasswordKey string

const (
	BirthYear      PasswordKey = "byr"
	IssueYear                  = "iyr"
	ExpirationYear             = "eyr"
	Height                     = "hgt"
	HairColor                  = "hcl"
	EyeColor                   = "ecl"
	PassportID                 = "pid"
	CountryID                  = "cid"
)

var requiredKeys = [...]PasswordKey{
	BirthYear,
	IssueYear,
	ExpirationYear,
	Height,
	HairColor,
	EyeColor,
	PassportID,
	// "CountryID",
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

func createPasswordMap(passwordRaw string) map[PasswordKey]string {
	password := map[PasswordKey]string{}

	keyValues := strings.Fields(passwordRaw)

	for _, keyValue := range keyValues {
		split := strings.Split(keyValue, ":")
		checkIndexArray(split, 1)
		password[PasswordKey(split[0])] = split[1]
	}

	return password
}

func transformToPasswordMap(passwordsRaw []string) []map[PasswordKey]string {
	var passwords []map[PasswordKey]string

	for _, str := range passwordsRaw {
		if str == "" {
			continue
		}

		password := createPasswordMap(str)
		passwords = append(passwords, password)
	}

	return passwords
}

func prepareData() []map[PasswordKey]string {
	data := loadFileContent("input.txt")
	passwordsRaw := strings.Split(data, "\n\n")

	return transformToPasswordMap(passwordsRaw)
}

func isPasswordValid(password map[PasswordKey]string) bool {
	for _, key := range requiredKeys {
		if _, ok := password[key]; !ok {
			return false
		}
	}

	return true
}

func getValidPasswordsCount(passwords []map[PasswordKey]string) int {
	count := 0

	for _, password := range passwords {
		if isPasswordValid(password) {
			count = count + 1
		}
	}

	return count
}

func isPasswordValidString(password map[PasswordKey]string) bool {
	if !isPasswordValid(password) {
		return false
	}

	birthday, error := strconv.Atoi(password[BirthYear])
	check(error)
	if birthday < 1920 || birthday > 2002 {
		return false
	}

	issue, error := strconv.Atoi(password[IssueYear])
	check(error)
	if issue < 2010 || issue > 2020 {
		return false
	}

	expiration, error := strconv.Atoi(password[ExpirationYear])
	check(error)
	if expiration < 2020 || expiration > 2030 {
		return false
	}

	heightReg, error := regexp.Compile(`^(\d+)(cm|in)$`)
	check(error)
	height := heightReg.MatchString(password[Height])
	if !height {
		return false
	}
	groups := heightReg.FindStringSubmatch(password[Height])
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

	hair, error := regexp.MatchString(`^#[0-9a-f]{6}$`, password[HairColor])
	check(error)
	if !hair {
		return false
	}

	eye, error := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, password[EyeColor])
	check(error)
	if !eye {
		return false
	}

	id, error := regexp.MatchString(`^\d{9}$`, password[PassportID])
	check(error)
	if !id {
		return false
	}

	return true
}

func getValidPasswordsCountStrinct(passwords []map[PasswordKey]string) int {
	count := 0

	for _, password := range passwords {
		if isPasswordValidString(password) {
			count = count + 1
		}
	}

	return count
}

func main() {
	passwords := prepareData()

	fmt.Println(getValidPasswordsCount(passwords))
	fmt.Println(getValidPasswordsCountStrinct(passwords))
}
