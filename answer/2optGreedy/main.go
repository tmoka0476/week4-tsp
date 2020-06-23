package main

import (
	"flag"
	"fmt"
	"os"

	"../dist"
	"../readCsv"
)

func main() {
	var (
		inputfile = flag.String("i", "", "input file")
		n         = flag.Int("n", 0, "the number of data")
		outfile   = flag.String("o", "", "output file")
	)
	flag.Parse()
	N := *n

	p := make([][3]float64, N)

	p = readCsv.Read(*inputfile, p)

	// find the Nearest Neighbor for each point
	sumDist := 0.0
	for i := 0; i < N-1; i++ {
		_, d := dist.FindNearest(p, i, N)
		sumDist = sumDist + d
	}

	// do 2-opt improvement
	for i := 0; i < 500; i++ {
		for j := 0; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				dist.Opt2(p, j, k, N)
			}
		}
	}

	f, err := os.OpenFile(*outfile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintln(f, "index")
	for i, _ := range p {
		fmt.Fprintln(f, p[i][0])

	}
}
