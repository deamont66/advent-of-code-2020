package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
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

type actionTypes = string

const (
	north   actionTypes = "N"
	south   actionTypes = "S"
	east    actionTypes = "E"
	west    actionTypes = "W"
	left    actionTypes = "L"
	right   actionTypes = "R"
	forward actionTypes = "F"
)

type action struct {
	actionType actionTypes
	parameter  int
}

func newAction(actionRow string) action {
	actionRow = strings.TrimSpace(actionRow)

	actionType := actionRow[:1]
	parameter, err := strconv.Atoi(actionRow[1:])
	check(err)

	return action{actionType: actionType, parameter: parameter}
}

type ship struct {
	x        int
	y        int
	rotation int
}

func normalizeAngle(angle int) int {
	return ((angle % 360) + 360) % 360
}

func (s ship) apply(a action) ship {
	// fmt.Println("before", s, a)
	if a.actionType == forward {
		if s.rotation == 0 {
			a.actionType = east
		} else if s.rotation == 90 {
			a.actionType = south
		} else if s.rotation == 180 {
			a.actionType = west
		} else if s.rotation == 270 {
			a.actionType = north
		}
	}

	switch a.actionType {
	case north:
		s.y = s.y - a.parameter
	case south:
		s.y = s.y + a.parameter
	case east:
		s.x = s.x + a.parameter
	case west:
		s.x = s.x - a.parameter
	case right:
		s.rotation = normalizeAngle(s.rotation + a.parameter)
	case left:
		s.rotation = normalizeAngle(s.rotation - a.parameter)
	}

	// fmt.Println("after", s)

	return s
}

func (s ship) applyMultiple(arr []action) ship {
	for _, a := range arr {
		s = s.apply(a)
	}

	return s
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func (s ship) distanceFromStart() int {
	return abs(s.x) + abs(s.y)
}

type waypointShip struct {
	x        int
	y        int
	waypoint ship
}

func newWaypointShip() waypointShip {
	return waypointShip{waypoint: ship{x: 10, y: -1}}
}

func toRad(deg int) float64 {
	return float64(deg) * (math.Pi / 180.0)
}

func (w waypointShip) rotateWaypoint(angle int) waypointShip {
	sin := math.Sin(toRad(angle))
	cos := math.Cos(toRad(angle))

	p := w.waypoint
	xnew := float64(p.x)*cos - float64(p.y)*sin
	ynew := float64(p.x)*sin + float64(p.y)*cos

	p.x = int(math.Round(xnew))
	p.y = int(math.Round(ynew))

	w.waypoint = p
	return w
}

func (w waypointShip) apply(a action) waypointShip {
	switch a.actionType {
	case north, east, south, west:
		w.waypoint = w.waypoint.apply(a)
	case forward:
		w.x = w.x + w.waypoint.x*a.parameter
		w.y = w.y + w.waypoint.y*a.parameter
	case left:
		w = w.rotateWaypoint(-a.parameter)
	case right:
		w = w.rotateWaypoint(a.parameter)
	}
	return w
}

func (w waypointShip) applyMultiple(arr []action) waypointShip {
	for _, a := range arr {
		w = w.apply(a)
	}

	return w
}

func (w waypointShip) distanceFromStart() int {
	return abs(w.x) + abs(w.y)
}

func trasformToActions(actionsRaw []string) []action {
	actions := []action{}

	for _, actionRaw := range actionsRaw {
		actions = append(actions, newAction(actionRaw))
	}

	return actions
}

func prepareData(filename string) []action {
	fileContent := loadFileContent(filename)
	fileContent = strings.TrimSpace(fileContent)

	mapRaw := strings.Split(fileContent, "\n")
	planMap := trasformToActions(mapRaw)

	return planMap
}

func run(filename string) {
	actions := prepareData(filename)
	s := ship{}
	s = s.applyMultiple(actions)
	fmt.Println("distance from start", filename, s.distanceFromStart(), s)

	w := newWaypointShip()
	w = w.applyMultiple(actions)
	fmt.Println("distance from start (waypoint)", filename, w.distanceFromStart(), w)
}

func main() {
	run("example.txt")
	run("input.txt")
}
