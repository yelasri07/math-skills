package mathskills

import (
	"math"
	"sort"
)

func Average(nbs []int) float64 {
	if len(nbs) == 0 {
		return 0
	}

	somme := 0.0
	for i := 0; i < len(nbs); i++ {
		somme += float64(nbs[i])
	}

	return somme / float64(len(nbs))
}

func Median(nbs []int) float64 {
	if len(nbs) == 0 {
		return 0
	}

	sort.Ints(nbs)
	if len(nbs)%2 == 0 {
		return (float64(nbs[len(nbs)/2-1]) + float64(nbs[len(nbs)/2])) / 2
	} else {
		return float64(nbs[len(nbs)/2])
	}
}

func Variance(nbs []int) float64 {
	if len(nbs) == 0 {
		return 0
	}

	average := Average(nbs)
	variance := 0.0
	for i := 0; i < len(nbs); i++ {
		pow := float64(nbs[i]) - average
		variance += pow * pow
	}

	return variance / float64(len(nbs))
}

func StandardDeviation(nbs []int) float64 {
	variance := Variance(nbs)
	return math.Sqrt(variance)
}
