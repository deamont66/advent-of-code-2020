package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type flightGroup struct {
	allAnswers    []string
	commonAnswers []string
}

func insert(ss []string, s string) []string {
	i := sort.SearchStrings(ss, s)
	ss = append(ss, "")
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func copyOnlyCommonAnswers(commonAnswers []string, currentAnswers []string) []string {
	var newCommon []string

	for _, commonAnswer := range commonAnswers {
		if contains(currentAnswers, commonAnswer) {
			newCommon = insert(newCommon, commonAnswer)
		}
	}
	return newCommon
}

func newFlightGroup(groupAnswersRaw []string) flightGroup {
	var commonAnswers []string
	var allAnswers []string

	for i, answers := range groupAnswersRaw {
		currentAnswers := strings.Split(answers, "")
		sort.Strings(currentAnswers)

		for _, question := range currentAnswers {
			if !contains(allAnswers, question) {
				allAnswers = insert(allAnswers, question)
			}
		}

		if i == 0 {
			// (._.) commonAnswers = allAnswers does not copy underlying data by value
			// this results in them getting rewritten in next insert(allAnswers, ...) call
			commonAnswers = make([]string, len(allAnswers))
			copy(commonAnswers, allAnswers)
		} else {
			commonAnswers = copyOnlyCommonAnswers(commonAnswers, currentAnswers)
		}
	}

	return flightGroup{allAnswers: allAnswers, commonAnswers: commonAnswers}
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

func transformToFlightGroups(groupsRaw []string) []flightGroup {
	var groups []flightGroup

	for _, groupRaw := range groupsRaw {
		if groupRaw == "" {
			continue
		}
		groupRaw = strings.TrimSpace(groupRaw)
		groupAnswersRaw := strings.Split(groupRaw, "\n")
		groups = append(groups, newFlightGroup(groupAnswersRaw))
	}

	return groups
}

func prepareTestData() []flightGroup {
	data := loadFileContent("example.txt")
	groups := strings.Split(data, "\n\n")

	return transformToFlightGroups(groups)
}

func prepareData() []flightGroup {
	data := loadFileContent("input.txt")
	groups := strings.Split(data, "\n\n")

	return transformToFlightGroups(groups)
}

func getNumberOfCommonAnswares(groups []flightGroup) int {
	var sum int
	for _, group := range groups {
		sum = sum + len(group.commonAnswers)
	}

	return sum
}

func getNumberOfAllAnswares(groups []flightGroup) int {
	var sum int
	for _, group := range groups {
		sum = sum + len(group.allAnswers)
	}

	return sum
}

func main() {
	flightGroups := prepareTestData()

	fmt.Println("sum of all answers (example)", getNumberOfAllAnswares(flightGroups))
	fmt.Println("sum of common answers (example)", getNumberOfCommonAnswares(flightGroups))

	flightGroups = prepareData()
	fmt.Println("sum of all answers", getNumberOfAllAnswares(flightGroups))
	fmt.Println("sum of common answers", getNumberOfCommonAnswares(flightGroups))
}
