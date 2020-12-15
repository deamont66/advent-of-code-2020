package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"sync"
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

type bus struct {
	busNumber int
}

func (b bus) nextDeparture(time int) int {
	if time%b.busNumber == 0 {
		return time
	}

	return time/b.busNumber*b.busNumber + b.busNumber
}

func prepareData(filename string) (int, []bus) {
	fileContent := loadFileContent(filename)
	fileContent = strings.TrimSpace(fileContent)

	rows := strings.Split(fileContent, "\n")
	timeS := strings.TrimSpace(rows[0])
	time, err := strconv.Atoi(timeS)
	check(err)
	busesRaw := strings.Split(rows[1], ",")
	buses := []bus{}

	for _, busNumberS := range busesRaw {
		if busNumberS == "x" {
			buses = append(buses, bus{})
			continue
		}
		busNumber, err := strconv.Atoi(busNumberS)
		check(err)
		buses = append(buses, bus{busNumber: busNumber})
	}

	return time, buses
}

func nextBus(time int, buses []bus) (int, bus) {
	nextDeparture := math.MaxInt64
	nextBus := bus{}

	for _, b := range buses {
		if b.busNumber == 0 {
			continue
		}
		busNextDeparture := b.nextDeparture(time)
		if busNextDeparture < nextDeparture {
			nextDeparture = busNextDeparture
			nextBus = b
		}
	}

	return nextDeparture, nextBus
}

func matchRequestedBusPlan(buses []bus, start int) bool {
	if len(buses) == 0 {
		return true
	}
	for i, b := range buses {
		if b.busNumber == 0 {
			continue
		}
		if (start+i)%b.busNumber != 0 {
			return false
		}
	}

	return true
}

func getBiggestBusNumber(buses []bus) (int, int) {
	max := 0
	maxI := 0
	for i, b := range buses {
		if b.busNumber > max {
			max = b.busNumber
			maxI = i
		}
	}

	return maxI, max
}

func getEarliestBusPlanTimestamp(buses []bus, start int, stop int) (int, error) {
	offset, biggestBusNumber := getBiggestBusNumber(buses)
	start = (start / biggestBusNumber) * biggestBusNumber

	for t := start; stop == 0 || t < stop; t += biggestBusNumber {
		if matchRequestedBusPlan(buses, t-offset) {
			return t - offset, nil
		}
	}

	return 0, errors.New("could not found earliest timestamp")
}

func run(filename string, start int, stop int, wg *sync.WaitGroup) {
	fmt.Println("processing", filename)
	time, buses := prepareData(filename)

	dept, b := nextBus(time, buses)
	waitTime := dept - time
	fmt.Println("next bus", b, waitTime, waitTime*b.busNumber)

	t, _ := getEarliestBusPlanTimestamp(buses, start, stop)
	if t != 0 {
		fmt.Println("!!!earliest requested start time", t)
	}

	if wg != nil {
		wg.Done()
	}
}

func main() {
	run("example.txt", 0, 0, nil)
	run("example2.txt", 0, 0, nil)
	run("example3.txt", 0, 0, nil)
	run("example4.txt", 0, 0, nil)
	run("example5.txt", 0, 0, nil)
	run("example6.txt", 0, 0, nil)

	var wg sync.WaitGroup
	for batch := 1; batch <= 100; batch++ {
		// it was in the 5th batch :)
		for i := 0; i < 20; i++ {
			wg.Add(1)
			go run("input.txt", batch*100000000000000+i*5000000000000, batch*100000000000000+(i+1)*5000000000000, &wg)
		}
		wg.Wait()
		fmt.Println(batch, "batch finished")
	}
}
