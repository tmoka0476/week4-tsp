package readCsv

import (
	"encoding/csv"
	"os"
	"strconv"
)

func Read(filename string, m [][3]float64) [][3]float64 {
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
		m[i] = [3]float64{float64(i), f0, f1}
	}

	return m
}
