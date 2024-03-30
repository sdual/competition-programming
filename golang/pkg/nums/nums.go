package nums

import (
	"math"
)

func Max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	result := nums[0]
	for i := 0; i < len(nums); i++ {
		result = int(math.Max(float64(result), float64(nums[i])))
	}
	return result
}

func Min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	result := nums[0]
	for i := 0; i < len(nums); i++ {
		result = int(math.Min(float64(result), float64(nums[i])))
	}
	return result
}
