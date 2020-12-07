package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type bagRule struct {
	numberOfBags int
	bagName      string
}

type bag struct {
	rules []bagRule
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

func parseBagInfo(bagString string) (int, string) {
	bagString = strings.TrimSpace(bagString)
	bagRegexp := regexp.MustCompile("^(\\d*) ?([a-z]+ [a-z]+) +bags?$")

	nameRegexpResult := bagRegexp.FindStringSubmatch(bagString)
	checkIndexArray(nameRegexpResult, 2)

	countString := strings.TrimSpace(nameRegexpResult[1])
	if countString == "" {
		countString = "0"
	}
	count, err := strconv.Atoi(countString)
	check(err)
	name := strings.TrimSpace(nameRegexpResult[2])

	return count, name
}

func trasformToBagMap(rules []string) map[string]bag {
	bags := map[string]bag{}

	for _, rule := range rules {
		parts := strings.SplitN(rule, " contain ", 2)
		checkIndexArray(parts, 1)

		_, mainBagName := parseBagInfo(parts[0])

		parts[1] = strings.TrimRight(parts[1], ".")
		bagRulesRaw := strings.Split(parts[1], ",")

		var bagRules []bagRule
		for _, bagRuleRaw := range bagRulesRaw {
			count, name := parseBagInfo(bagRuleRaw)

			if count != 0 { // no other
				currentRule := bagRule{bagName: name, numberOfBags: count}
				bagRules = append(bagRules, currentRule)
			}
		}
		bags[mainBagName] = bag{rules: bagRules}
	}

	return bags
}

func contains(bags map[string]bag, singleBagName string, needle string) bool {
	singleBag := bags[singleBagName]

	for _, singleBagRule := range singleBag.rules {
		if singleBagRule.bagName == needle { // we found him
			return true
		} else if contains(bags, singleBagRule.bagName, needle) { // we need to go deeper
			return true
		}
	}

	return false
}

func findAllContainingBag(bags map[string]bag, needle string) int {
	count := 0
	for bagName := range bags {
		if bagName == needle {
			continue
		}
		if contains(bags, bagName, needle) {
			count = count + 1
		}
	}

	return count
}

func findNumberOfNestedBags(bags map[string]bag, name string) int {
	count := 1
	for _, nestedBagRule := range bags[name].rules {
		count = count + nestedBagRule.numberOfBags*findNumberOfNestedBags(bags, nestedBagRule.bagName)
	}

	return count
}

func prepareData(filename string) map[string]bag {
	fileContent := loadFileContent(filename)
	fileContent = strings.TrimSpace(fileContent)

	rules := strings.Split(fileContent, "\n")
	bagMap := trasformToBagMap(rules)

	return bagMap
}

func main() {
	bags := prepareData("example.txt")
	fmt.Println("bags containing shiny gold bag (example)", findAllContainingBag(bags, "shiny gold"))
	fmt.Println("nested bags inside of shiny gold bag (example)", findNumberOfNestedBags(bags, "shiny gold")-1)

	bags = prepareData("input.txt")
	fmt.Println("bags containing shiny gold bag", findAllContainingBag(bags, "shiny gold"))
	fmt.Println("nested bags inside of shiny gold bag", findNumberOfNestedBags(bags, "shiny gold")-1)
}
