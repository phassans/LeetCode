package easy

import "math"

func relativeSortArray(A, B []int) []int {
	count := [math.MaxInt8]int{}
	for _, a := range A {
		count[a]++
	}

	res := make([]int, 0, len(A))
	for _, b := range B {
		for count[b] > 0 {
			res = append(res, b)
			count[b]--
		}
	}
	for i := 0; i < math.MaxInt8; i++ {
		for count[i] > 0 {
			res = append(res, i)
			count[i]--
		}
	}

	return res
}

func relativeSortArrayMap(A, B []int) []int {
	countA := make(map[int]int)
	for _, a := range A {
		if _, ok := countA[a]; ok {
			countA[a]++
		} else {
			countA[a] = 1
		}
	}

	var res []int
	for _, b := range B {
		if _, ok := countA[b]; ok {
			for countA[b] > 0 {
				res = append(res, b)
				countA[b]--
			}
		}
	}

	for key, _ := range countA {
		for countA[key] > 0 {
			res = append(res, key)
			countA[key]--
		}
	}

	return res
}
