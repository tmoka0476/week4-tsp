package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
)

const N = 64

func main() {
	m := make([][3]float64, N)

	file, err := os.Open("../../input_3.csv")
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
		m[i] = [3]float64{float64(i), f0, f1}
	}
	sumDist := 0.0
	for i := 0; i < N-1; i++ {
		_, d := findNearest(m, i)
		sumDist = sumDist + d
	}
	fmt.Println("index")
	for i, _ := range m {
		fmt.Println(m[i][0])
	}
}

func dist(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(dx*dx + dy*dy)
}

func findNearest(p [][3]float64, i int) ([][3]float64, float64) {
	minDist := 1000000.0
	index := 0
	for j := i + 1; j < N; j++ {
		d := dist(p[i][1], p[i][2], p[j][1], p[j][2])
		if minDist > d {
			minDist = d
			index = j
		}
	}
	p[i+1], p[index] = p[index], p[i+1]
	return p, minDist
}
