package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	Id   int
	X    float64
	Y    float64
	Next *Point
}

func main() {
	var (
		inputfile = flag.String("i", "", "input file")
		n         = flag.Int("n", 0, "the number of data")
		//outfile   = flag.String("o", "", "output file")
	)
	flag.Parse()
	N := *n

	points := []Point{}
	points = readCSV(*inputfile, points)

	//fmt.Println(points, N)
	pp := findNearest(points, N)
	for i, _ := range pp {
		a := pp[i].Next
		fmt.Println(*a)
	}
}

func readCSV(filename string, points []Point) []Point {
	file, err := os.Open("../../" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string
	line, err = reader.Read()

	for i := 0; ; i++ {
		line, err = reader.Read()
		if err != nil {
			break
		}
		f0, _ := strconv.ParseFloat(line[0], 64)
		f1, _ := strconv.ParseFloat(line[1], 64)
		p := &Point{Id: i, X: f0, Y: f1}
		points = append(points, *p)
	}
	return points
}

func dist(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func findNearest(points []Point, N int) []Point {
	visited := make([]bool, N)
	visited[0] = true
	k := 0
	for i := 0; i < N; i++ {
		minDist := math.Inf(0)
		for j := 0; j < N; j++ {
			if !visited[j] {
				d := dist(points[k], points[j])
				if minDist > d {
					minDist = d
					points[k].Next = &points[j]
				}
			}
		}
		visited[points[k].Next.Id] = true
		k = points[k].Next.Id
	}
	return points
}
