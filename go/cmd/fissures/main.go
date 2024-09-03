package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`
}

type Point struct {
	x int
	y int
}

type Line struct {
	p1 *Point
	p2 *Point
}

func isHOrV(p1 Point, p2 Point) bool {
	return p1.x == p2.x || p1.y == p2.y
}

func parsePoint(line string) (*Point, error) {
	points := strings.Split(line, ",")

	px, err := strconv.Atoi(points[0])
	if err != nil {
		return nil, err
	}

	py, err := strconv.Atoi(points[1])
	if err != nil {
		return nil, err
	}

	return &Point{
		x: px,
		y: py,
	}, nil
}

func parseLine(line string) (*Line, error) {
	points := strings.Split(line, " -> ")

	p1, err := parsePoint(points[0])
	if err != nil {
		return nil, err
	}

	p2, err := parsePoint(points[1])
	if err != nil {
		return nil, err
	}

	return &Line{
		p1,
		p2,
	}, nil
}

func main() {
	lines := []Line{}
	for _, l := range strings.Split(getInput(), "\n") {
		line, err := parseLine(l)
		if err != nil {
			log.Fatal("Can't parse line")
		}

		lines = append(lines, *line)
	}

	fmt.Printf("%+v", lines)

}
