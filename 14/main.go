package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
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

type bitMask struct {
	clearMask uint64 // 1 bit leaves bit unchanged
	setMask   uint64 // 0 bit leaves bit unchanged
}

func newBitMask() bitMask {
	clearMask := uint64(math.MaxUint64)
	setMask := uint64(0)

	return bitMask{setMask: setMask, clearMask: clearMask}
}

func newBitMaskFromString(maskString string) bitMask {
	maskString = strings.TrimSpace(maskString)

	clearMask := uint64(math.MaxUint64)
	setMask := uint64(0)

	for _, ch := range maskString {
		clearMask = clearMask << 1
		setMask = setMask << 1
		switch ch {
		case '1':
			{
				clearMask = clearMask | 1
				setMask = setMask | 1
			}
		case '0':
			{
				clearMask = clearMask & ^uint64(1)
				setMask = setMask & ^uint64(1)
			}
		case 'X':
			{
				clearMask = clearMask | 1
				setMask = setMask & ^uint64(1)
			}
		}
	}

	return bitMask{setMask: setMask, clearMask: clearMask}
}

func (mask bitMask) apply(number uint64) uint64 {
	result := number
	result = result | mask.setMask
	result = result & mask.clearMask
	return result
}

type memory struct {
	mask bitMask
	data map[uint64]uint64
}

func newMemory() memory {
	return memory{mask: newBitMask(), data: map[uint64]uint64{}}
}

func (m *memory) setMask(mask bitMask) {
	m.mask = mask
}

func (m *memory) set(addr uint64, value uint64) uint64 {
	maskedValue := m.mask.apply(value)

	m.data[addr] = maskedValue

	return maskedValue
}

func (m *memory) setDec2(addr uint64, value uint64) uint64 {
	// set 1 from mask
	maskedAddr := addr | m.mask.setMask

	clear := (m.mask.clearMask & (^uint64(0) >> 28))
	diff := clear ^ m.mask.setMask
	diffs := []uint64{0}

	diff2 := diff
	for i := 0; diff2 != 0; i++ {
		if diff2&1 == 1 {
			for _, d := range diffs {
				diffs = append(diffs, d|1<<i)
			}
		}
		diff2 = diff2 >> 1
	}

	// clear floating bits
	maskedAddr = maskedAddr & ^diff
	for _, d := range diffs {
		realAddr := maskedAddr | d
		m.data[realAddr] = value
	}

	return maskedAddr
}

func (m memory) sumSetValues() uint64 {
	sum := uint64(0)
	for _, value := range m.data {
		sum = sum + value
	}

	return sum
}

func processInstructions(instructions []string) (memory, memory) {
	m := newMemory()
	m2 := newMemory()

	maskReg := regexp.MustCompile("^mask = ([X10]{36})$")
	assignReg := regexp.MustCompile("^mem\\[(\\d{1,36})\\] = (\\d{1,36})$")
	for _, instruction := range instructions {
		instruction = strings.TrimSpace(instruction)
		match := maskReg.FindStringSubmatch(instruction)
		if match != nil {
			m.setMask(newBitMaskFromString(match[1]))
			m2.setMask(newBitMaskFromString(match[1]))
			continue
		}
		match = assignReg.FindStringSubmatch(instruction)
		if match != nil {
			addr, _ := strconv.ParseUint(match[1], 10, 36)
			val, _ := strconv.ParseUint(match[2], 10, 36)
			m.set(addr, val)
			m2.setDec2(addr, val)
			continue
		}

		panic("instruction \"" + instruction + "\" cannot be parsed")
	}

	return m, m2
}

func run(filename string) {
	data := loadFileContent(filename)
	data = strings.TrimSpace(data)
	instructions := strings.Split(data, "\n")

	m, m2 := processInstructions(instructions)

	fmt.Println("sum", filename, m.sumSetValues())
	fmt.Println("sum2", filename, m2.sumSetValues())
}

func main() {
	// run("example.txt") example from first part cannot be run with rules from the second (too many floating bits)
	run("example2.txt")
	run("input.txt")
}
