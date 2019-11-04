package easy

import (
	"math"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	maxDiff := math.MaxInt32
	sort.Ints(arr)

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < maxDiff {
			maxDiff = arr[i] - arr[i-1]
		}
		if maxDiff == 1 {
			break
		}
	}

	var res [][]int

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] == maxDiff {
			temp := []int{arr[i], arr[i-1]}
			res = append(res, temp)
		}
	}

	return res
}
