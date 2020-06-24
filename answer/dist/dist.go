package dist

import (
	"math"
)

func Dist(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(dx*dx + dy*dy)
}

func FindNearest(p [][3]float64, i int, N int) ([][3]float64, float64) {
	minDist := math.Inf(0)
	index := 0
	for j := i + 1; j < N; j++ {
		d := Dist(p[i][1], p[i][2], p[j][1], p[j][2])
		if minDist > d {
			minDist = d
			index = j
		}
	}
	p[i+1], p[index] = p[index], p[i+1]
	return p, minDist
}

func Opt2(p [][3]float64, i, j, N int) {
	var d_bf, d_af float64
	if j == (N - 1) {
		d_bf = Dist(p[i][1], p[i][2], p[i+1][1], p[i+1][2]) + Dist(p[j][1], p[j][2], p[0][1], p[0][2])
		d_af = Dist(p[i][1], p[i][2], p[j][1], p[j][2]) + Dist(p[0][1], p[0][2], p[i+1][1], p[i+1][2])
	} else {
		d_bf = Dist(p[i][1], p[i][2], p[i+1][1], p[i+1][2]) + Dist(p[j][1], p[j][2], p[j+1][1], p[j+1][2])
		d_af = Dist(p[i][1], p[i][2], p[j][1], p[j][2]) + Dist(p[j+1][1], p[j+1][2], p[i+1][1], p[i+1][2])
	}
	if d_bf > d_af {
		n := i + 1
		m := j
		for {
			p[n], p[m] = p[m], p[n]
			n++
			m--
			if n >= m {
				break
			}
		}
	}
}
