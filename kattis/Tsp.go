package kattis

import (
	"fmt"
	"math"
	"time"
)

func Tsp() {
	start := time.Now()
	fPoints := ReadMatrix()
	dist := DistFromFPoints(fPoints)
	tour := NearestNeighbourFromDist(dist)
	cost := CostOfClosedTour(dist, tour)
	fmt.Println("nn tour, cost=", cost, ":", tour)
	tour = ThreeOpt(tour, dist, start)
	cost = CostOfClosedTour(dist, tour)
	fmt.Println("three opt tour, cost=", cost,  ":", tour)
}

func ReadMatrix() [][]float64 {
	N := 0
	fmt.Scan(&N)
	result := make([][]float64, 0, N)
	for i := 0; i < N; i++ {
		point := make([]float64, 2)
		fmt.Scan(&point[0])
		fmt.Scan(&point[1])
		result = append(result, point)
	}
	return result
}

func DistFromFPoints(fpoints [][]float64) [][]int {
	N := len(fpoints)
	g := make([][]int, N)
	for i, p := range fpoints {
		g[i] = make([]int, N)
		for j := 0; j < i; j++ {
			q := fpoints[j]
			d := RoundedEuclideanDistFloat64(p, q)
			g[i][j] = int(d)
			g[j][i] = int(d)
		}
	}
	return g
}

func RoundedEuclideanDistFloat64(a []float64, b []float64) int {
	return int(math.Round(math.Hypot(a[0]-b[0], a[1]-b[1])))
}

func CostOfClosedTour(g [][]int, tour []int) int {
	result := 0
	current := tour[0]
	for i := 1; i < len(tour); i++ {
		next := tour[i]
		result += g[current][next]
		current = next
	}
	result += g[0][current]
	return result
}

func NearestNeighbourFromDist(dist [][]int) []int {
	N := len(dist)
	tour := make([]int, 0, N)
	visited := make([]bool, N)
	current := 0
	visited[current] = true
	tour = append(tour, current)
	for len(tour) < N {
		next, minDist := current, math.MaxInt64
		for i := range dist {
			d := dist[current][i]
			if !visited[i] && d < minDist {
				next, minDist = i, d
			}
		}
		tour = append(tour, next)
		visited[next] = true
		current = next
	}
	return tour
}

// 3-opt adapted from https://en.wikipedia.org/wiki/3-opt

func ReverseSegmentIfBetter(tour []int, dist [][]int, i, j, k int) int {
	N := len(tour)
	A, B, C, D, E, F := tour[(i+N-1)%N], tour[i], tour[j-1], tour[j], tour[k-1], tour[k]
	d0 := dist[A][B] + dist[C][D] + dist[E][F]
	d1 := dist[A][C] + dist[B][D] + dist[E][F]
	d2 := dist[A][B] + dist[C][E] + dist[D][F]
	d3 := dist[A][D] + dist[E][B] + dist[C][F]
	d4 := dist[F][B] + dist[C][D] + dist[E][A]

	// If reversing tour[i:j] would make the tour shorter, then do it.
	if d0 > d1 {
		ReverseSliceInts(tour, i, j)
		return -d0 + d1
	}
	if d0 > d2 {
		ReverseSliceInts(tour, j, k)
		return -d0 + d2
	}
	if d0 > d4 {
		ReverseSliceInts(tour, i, k)
		ReverseSliceInts(tour, j, k)
		return -d0 + d4
	}
	if d0 > d3 {
		tmp := CopyInts(tour[j:k])
		tmp = append(tmp, tour[i:j]...)
		copy(tour[i:k], tmp)
		ReverseSliceInts(tour, j, k)
		return -d0 + d3
	}
	return 0
}

func ThreeOpt(tour []int, dist [][]int, start time.Time) []int {
	N := len(tour)
	allSegs := make([][]int, 0, N*N)
	for i := 0; i < N; i++ {
		for j := i + 2; j < MinInt(N,100); j++ {
			for k := j + 2; k < MinInt(N,100); k++ {
				allSegs = append(allSegs, []int{i, j, k})
			}
		}
	}
	fmt.Println("time:", time.Since(start))
	for true {
		delta := 0
		for _, x := range allSegs {
			delta += ReverseSegmentIfBetter(tour, dist, x[0], x[1], x[2])
			if time.Since(start).Milliseconds() > 1865 {
				return tour
			}
		}
		if delta >= 0 {
			break
		}
	}
	return tour
}

// Library

func ReverseInts(ints []int) {
	for i, j := 0, len(ints)-1; i < j; i, j = i+1, j-1 {
		ints[i], ints[j] = ints[j], ints[i]
	}
}

func ReverseSliceInts(ints []int, i, j int) {
	for k, l := i, j-1; k < l; k, l = k+1, l-1 {
		ints[k], ints[l] = ints[l], ints[k]
	}
}

func CopyInts(xs []int) []int {
	result := make([]int, len(xs))
	copy(result, xs)
	return result
}

func MinInt(x ...int) int {
	result := math.MaxInt64
	for _, num := range x {
		if num < result {
			result = num
		}
	}
	return result
}
