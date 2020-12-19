package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
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

type valueRange struct {
	min int
	max int
}

func newValueRange(min string, max string) valueRange {
	minNumber, err := strconv.Atoi(min)
	check(err)
	maxNumber, err := strconv.Atoi(max)
	check(err)

	return valueRange{min: minNumber, max: maxNumber}
}

func (v valueRange) isValid(number int) bool {
	return number >= v.min && number <= v.max
}

type rule struct {
	fieldName   string
	firstRange  valueRange
	secondRange valueRange
}

func newRule(ruleRaw string) rule {
	ruleRaw = strings.TrimSpace(ruleRaw)
	r := regexp.MustCompile("^([^:]*): (\\d+)-(\\d+) or (\\d+)-(\\d+)$")

	matches := r.FindStringSubmatch(ruleRaw)
	fieldName := matches[1]
	firstRange := newValueRange(matches[2], matches[3])
	secondRange := newValueRange(matches[4], matches[5])

	return rule{fieldName: fieldName, firstRange: firstRange, secondRange: secondRange}
}

func (r rule) isValid(number int) bool {
	return r.firstRange.isValid(number) || r.secondRange.isValid(number)
}

func parseRules(rawRules []string) []rule {
	rules := []rule{}

	for _, rawRule := range rawRules {
		rules = append(rules, newRule(rawRule))
	}

	return rules
}

type ticket struct {
	numbers []int
}

func newTicket(rawTicket string) ticket {
	rawTicket = strings.TrimSpace(rawTicket)

	rawNumbers := strings.Split(rawTicket, ",")
	numbers := []int{}
	for _, rawNumber := range rawNumbers {
		number, err := strconv.Atoi(rawNumber)
		check(err)
		numbers = append(numbers, number)
	}

	return ticket{numbers: numbers}
}

func (t ticket) isValid(rules []rule) bool {
	isValidTicket := true
	for _, n := range t.numbers {
		isValid := false
		for _, r := range rules {
			if r.isValid(n) {
				isValid = true
				break
			}
		}

		if !isValid {
			isValidTicket = false
			break
		}
	}
	return isValidTicket
}

func parseTickets(rawTickets []string) []ticket {
	tickets := []ticket{}

	for _, rawTicket := range rawTickets {
		tickets = append(tickets, newTicket(rawTicket))
	}

	return tickets
}

func sumInvalidNumbers(rules []rule, tickets []ticket) int {
	sum := 0
	for _, t := range tickets {
		for _, n := range t.numbers {
			isValid := false
			for _, r := range rules {
				if r.isValid(n) {
					isValid = true
					break
				}
			}
			if !isValid {
				sum = sum + n
			}
		}
	}

	return sum
}

func filterInvalidTickets(rules []rule, tickets []ticket) []ticket {
	validTickets := []ticket{}

	for _, t := range tickets {
		if t.isValid(rules) {
			validTickets = append(validTickets, t)
		}
	}

	return validTickets
}

func copyMap(m map[string]int) map[string]int {
	cm := map[string]int{}

	for k, v := range m {
		cm[k] = v
	}

	return cm
}

func copyMap2(m map[int]bool) map[int]bool {
	cm := map[int]bool{}

	for k, v := range m {
		cm[k] = v
	}

	return cm
}

func getMappingRec(acc map[string]int, usedIndexes map[int]bool, tickets []ticket, rules []rule) (map[string]int, error) {
	if len(rules) == 0 {
		return acc, nil
	}

	r := rules[0]
	for i := 0; i < len(tickets[0].numbers); i++ {
		if usedIndexes[i] {
			continue
		}

		isMatch := true
		for _, t := range tickets {
			if !r.isValid(t.numbers[i]) {
				isMatch = false
				break
			}
		}
		if isMatch {
			accCopy := copyMap(acc)
			accCopy[r.fieldName] = i
			usedIndexesCopy := copyMap2(usedIndexes)
			usedIndexesCopy[i] = true

			mapping, err := getMappingRec(accCopy, usedIndexesCopy, tickets, rules[1:])
			if err == nil {
				return mapping, nil
			}
		}
	}

	return nil, errors.New("not found")
}

func getMapping(tickets []ticket, rules []rule) (map[string]int, error) {
	acc := map[string]int{}
	usedIndexes := map[int]bool{}

	return getMappingRec(acc, usedIndexes, tickets, rules)
}

func multiplyDepartureValue(mapping map[string]int, t ticket) int {
	multiply := 1

	for fieldName, i := range mapping {
		if strings.HasPrefix(fieldName, "departure") {
			multiply = multiply * t.numbers[i]
		}
	}

	return multiply
}

func run(filename string) {
	data := loadFileContent(filename)
	data = strings.TrimSpace(data)
	inputParts := strings.Split(data, "\n\n")

	rawRules := strings.Split(strings.TrimSpace(inputParts[0]), "\n")
	rules := parseRules(rawRules)

	rawTickets := strings.Split(strings.TrimSpace(inputParts[1]), "\n")[1:]
	myTicket := parseTickets(rawTickets)[0]
	rawTickets = strings.Split(strings.TrimSpace(inputParts[2]), "\n")[1:]
	nearbyTickets := parseTickets(rawTickets)

	fmt.Println("sum of invalid ticket numbers", filename, sumInvalidNumbers(rules, nearbyTickets))

	validNearbyTickets := filterInvalidTickets(rules, nearbyTickets)
	fieldMapping, _ := getMapping(validNearbyTickets, rules)

	fmt.Println("departure values multiply", filename, multiplyDepartureValue(fieldMapping, myTicket))
}

func main() {
	run("example.txt")
	run("example2.txt")
	run("input.txt")
}
