package main

import (
	"fmt"
	"os"

	"../dist"
	"../readCsv"
)

const N = 2048

func main() {
	p := make([][3]float64, N)

	p = readCsv.Read("input_6.csv", p)

	sumDist := 0.0
	for i := 0; i < N-1; i++ {
		_, d := dist.FindNearest(p, i, N)
		sumDist = sumDist + d
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < N-1; j++ {
			for k := j + 1; k < N-1; k++ {
				dist.Exchange(p, j, k)
			}
		}
	}

	outfile, err := os.OpenFile("solution_greedy_6.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	fmt.Fprintln(outfile, "index")
	for i, _ := range p {
		fmt.Fprintln(outfile, p[i][0])
	}
}
