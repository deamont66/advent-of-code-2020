package main

import (
	"errors"
	"fmt"
	"io/ioutil"
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

type vertice4 struct {
	w int
	z int
	y int
	x int
}

func newVertice(x int, y int, z int, w int) vertice4 {
	return vertice4{x: x, y: y, z: z, w: w}
}

func (v vertice4) getHash() int {
	return v.w*100000000000 + v.z*100000000 + v.y*100000 + v.x
}

func (v vertice4) getNeighbours() []vertice4 {
	return []vertice4{
		newVertice(v.x+1, v.y+1, v.z+1, v.w+1),
		newVertice(v.x+1, v.y+1, v.z+1, v.w-1),
		newVertice(v.x+1, v.y+1, v.z+1, v.w),
		newVertice(v.x+1, v.y+1, v.z-1, v.w+1),
		newVertice(v.x+1, v.y+1, v.z-1, v.w-1),
		newVertice(v.x+1, v.y+1, v.z-1, v.w),
		newVertice(v.x+1, v.y+1, v.z, v.w+1),
		newVertice(v.x+1, v.y+1, v.z, v.w-1),
		newVertice(v.x+1, v.y+1, v.z, v.w),
		newVertice(v.x+1, v.y-1, v.z+1, v.w+1),
		newVertice(v.x+1, v.y-1, v.z+1, v.w-1),
		newVertice(v.x+1, v.y-1, v.z+1, v.w),
		newVertice(v.x+1, v.y-1, v.z-1, v.w+1),
		newVertice(v.x+1, v.y-1, v.z-1, v.w-1),
		newVertice(v.x+1, v.y-1, v.z-1, v.w),
		newVertice(v.x+1, v.y-1, v.z, v.w+1),
		newVertice(v.x+1, v.y-1, v.z, v.w-1),
		newVertice(v.x+1, v.y-1, v.z, v.w),
		newVertice(v.x+1, v.y, v.z+1, v.w+1),
		newVertice(v.x+1, v.y, v.z+1, v.w-1),
		newVertice(v.x+1, v.y, v.z+1, v.w),
		newVertice(v.x+1, v.y, v.z-1, v.w+1),
		newVertice(v.x+1, v.y, v.z-1, v.w-1),
		newVertice(v.x+1, v.y, v.z-1, v.w),
		newVertice(v.x+1, v.y, v.z, v.w+1),
		newVertice(v.x+1, v.y, v.z, v.w-1),
		newVertice(v.x+1, v.y, v.z, v.w),

		newVertice(v.x-1, v.y+1, v.z+1, v.w+1),
		newVertice(v.x-1, v.y+1, v.z+1, v.w-1),
		newVertice(v.x-1, v.y+1, v.z+1, v.w),
		newVertice(v.x-1, v.y+1, v.z-1, v.w+1),
		newVertice(v.x-1, v.y+1, v.z-1, v.w-1),
		newVertice(v.x-1, v.y+1, v.z-1, v.w),
		newVertice(v.x-1, v.y+1, v.z, v.w+1),
		newVertice(v.x-1, v.y+1, v.z, v.w-1),
		newVertice(v.x-1, v.y+1, v.z, v.w),
		newVertice(v.x-1, v.y-1, v.z+1, v.w+1),
		newVertice(v.x-1, v.y-1, v.z+1, v.w-1),
		newVertice(v.x-1, v.y-1, v.z+1, v.w),
		newVertice(v.x-1, v.y-1, v.z-1, v.w+1),
		newVertice(v.x-1, v.y-1, v.z-1, v.w-1),
		newVertice(v.x-1, v.y-1, v.z-1, v.w),
		newVertice(v.x-1, v.y-1, v.z, v.w+1),
		newVertice(v.x-1, v.y-1, v.z, v.w-1),
		newVertice(v.x-1, v.y-1, v.z, v.w),
		newVertice(v.x-1, v.y, v.z+1, v.w+1),
		newVertice(v.x-1, v.y, v.z+1, v.w-1),
		newVertice(v.x-1, v.y, v.z+1, v.w),
		newVertice(v.x-1, v.y, v.z-1, v.w+1),
		newVertice(v.x-1, v.y, v.z-1, v.w-1),
		newVertice(v.x-1, v.y, v.z-1, v.w),
		newVertice(v.x-1, v.y, v.z, v.w+1),
		newVertice(v.x-1, v.y, v.z, v.w-1),
		newVertice(v.x-1, v.y, v.z, v.w),

		newVertice(v.x, v.y+1, v.z+1, v.w+1),
		newVertice(v.x, v.y+1, v.z+1, v.w-1),
		newVertice(v.x, v.y+1, v.z+1, v.w),
		newVertice(v.x, v.y+1, v.z-1, v.w+1),
		newVertice(v.x, v.y+1, v.z-1, v.w-1),
		newVertice(v.x, v.y+1, v.z-1, v.w),
		newVertice(v.x, v.y+1, v.z, v.w+1),
		newVertice(v.x, v.y+1, v.z, v.w-1),
		newVertice(v.x, v.y+1, v.z, v.w),
		newVertice(v.x, v.y-1, v.z+1, v.w+1),
		newVertice(v.x, v.y-1, v.z+1, v.w-1),
		newVertice(v.x, v.y-1, v.z+1, v.w),
		newVertice(v.x, v.y-1, v.z-1, v.w+1),
		newVertice(v.x, v.y-1, v.z-1, v.w-1),
		newVertice(v.x, v.y-1, v.z-1, v.w),
		newVertice(v.x, v.y-1, v.z, v.w+1),
		newVertice(v.x, v.y-1, v.z, v.w-1),
		newVertice(v.x, v.y-1, v.z, v.w),
		newVertice(v.x, v.y, v.z+1, v.w+1),
		newVertice(v.x, v.y, v.z+1, v.w-1),
		newVertice(v.x, v.y, v.z+1, v.w),
		newVertice(v.x, v.y, v.z-1, v.w+1),
		newVertice(v.x, v.y, v.z-1, v.w-1),
		newVertice(v.x, v.y, v.z-1, v.w),
		newVertice(v.x, v.y, v.z, v.w+1),
		newVertice(v.x, v.y, v.z, v.w-1),
		// newVertice(v.x, v.y, v.z, v.w), skip middle point
	}
}

type vertice4Set struct {
	data map[int]vertice4
}

func newVertice4Set() vertice4Set {
	return vertice4Set{data: map[int]vertice4{}}
}

func (s *vertice4Set) add(v vertice4) bool {
	_, f := s.data[v.getHash()]
	if !f {
		s.data[v.getHash()] = v
	}
	return !f
}

func (s *vertice4Set) list() []vertice4 {
	res := []vertice4{}
	for _, v := range s.data {
		res = append(res, v)
	}
	return res
}

type pocketRow = map[int]bool
type pocketSubslice = map[int]pocketRow
type pocketSlice = map[int]pocketSubslice
type pocketData = map[int]pocketSlice

type pocket struct {
	data pocketData
}

func (p *pocket) initialize(ver vertice4) {
	if p.data == nil {
		p.data = pocketData{}
	}
	_, f := p.data[ver.w]
	if !f {
		p.data[ver.w] = pocketSlice{}
	}
	_, f = p.data[ver.w][ver.z]
	if !f {
		p.data[ver.w][ver.z] = pocketSubslice{}
	}
	_, f = p.data[ver.w][ver.z][ver.y]
	if !f {
		p.data[ver.w][ver.z][ver.y] = pocketRow{}
	}
}

func (p *pocket) set(ver vertice4, value bool) {
	p.initialize(ver)

	if value {
		p.data[ver.w][ver.z][ver.y][ver.x] = true
	} else {
		delete(p.data[ver.w][ver.z][ver.y], ver.x)
	}
}

func (p *pocket) get(ver vertice4) bool {
	p.initialize(ver)

	return p.data[ver.w][ver.z][ver.y][ver.x]
}

func (p pocket) activePoints() int {
	cnt := 0
	for _, sv := range p.data {
		for _, ssv := range sv {
			for _, rv := range ssv {
				for _, cv := range rv {
					if cv {
						cnt++
					}
				}
			}
		}
	}

	return cnt
}

func createPocketUniverse(rows []string) pocket {
	universe := pocket{}

	for ri, row := range rows {
		row = strings.TrimSpace(row)
		for i, col := range row {
			universe.set(newVertice(i, ri, 0, 0), col == '#')
		}
	}
	return universe
}

func copyPocketWithIndexesThatCouldChange(pocketUniverse pocket) (pocket, []vertice4) {
	result := pocket{}
	couldChange := newVertice4Set()
	for sk, sv := range pocketUniverse.data {
		for ssk, ssv := range sv {
			for rk, rv := range ssv {
				for ck, cv := range rv {
					currentVer := newVertice(ck, rk, ssk, sk)
					result.set(currentVer, cv)
					if cv {
						couldChange.add(currentVer)

						for _, n := range currentVer.getNeighbours() {
							couldChange.add(n)
						}
					}
				}
			}
		}
	}
	return result, couldChange.list()
}

func step(pocketUniverse pocket) pocket {
	nextPocket, couldChange := copyPocketWithIndexesThatCouldChange(pocketUniverse)

	for _, verticeForChange := range couldChange {
		activeNeighbours := 0
		for _, n := range verticeForChange.getNeighbours() {
			if pocketUniverse.get(n) {
				activeNeighbours++
			}
			if activeNeighbours > 3 {
				break
			}
		}

		if pocketUniverse.get(verticeForChange) {
			nextPocket.set(verticeForChange, activeNeighbours == 2 || activeNeighbours == 3)
		} else {
			nextPocket.set(verticeForChange, activeNeighbours == 3)
		}
	}
	return nextPocket
}

func printUniverse(pocketUniverse pocket) {
	for i, slice := range pocketUniverse.data {
		fmt.Println(i, slice)
	}
}

func run(filename string) {
	data := loadFileContent(filename)
	data = strings.TrimSpace(data)
	rows := strings.Split(data, "\n")

	pocketUniverse := createPocketUniverse(rows)
	for i := 0; i < 6; i++ {
		pocketUniverse = step(pocketUniverse)
	}

	fmt.Println("active after 6 steps", filename, pocketUniverse.activePoints())
}

func main() {
	run("example.txt")
	run("input.txt")
}
