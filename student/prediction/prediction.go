package prediction

import (
	"guess-it-2/student/prediction/calculations"
	"math"
)

func PredictRange(sliceX, sliceY []int) []int {
	if len(sliceX) == 1 {
		return []int{sliceY[0] - 9, sliceY[0] - 2}
	}
	result := []int{}
	a, b := calculations.CalculateLinearCoeff(sliceX, sliceY)
	y := a*float64(len(sliceX)) + b
	var min, max int

	min = int(math.Round(y)) - 45
	max = int(math.Round(y)) + 45

	result = append(result, min)
	result = append(result, max)

	return result
}
