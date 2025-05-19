package utils

import (
	"strconv"
	"sync"
)

// calculateBasePremium calculates the base premium based on population size and density
func CalculateBasePremium(population int64) float32 {
	switch {
	case population > 50000:
		return 100.0
	case population > 9938:
		return 70.0
	case population > 3000:
		return 50.0
	default:
		return 40.0
	}
}

// StrToFloat converts a string to float64
func StrToFloat(s string) float64 {
	newInt, _ := strconv.Atoi(s)
	return float64(newInt)
}

// ProcessLocations processes a slice of locations in batches using goroutines
func ProcessLocations[T any](locations []T, batchSize int, processor func(idx int)) {
	total := len(locations)
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		var wg sync.WaitGroup
		for j := i; j < end; j++ {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				processor(idx)
			}(j)
		}
		wg.Wait()
	}
}
