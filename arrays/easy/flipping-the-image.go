package easy

func flipAndInvertImage(A [][]int) [][]int {
	for k := 0; k < len(A); k++ {
		i, j := 0, len(A[k])-1
		for i < j {
			A[k][i], A[k][j] = (A[k][j]), (A[k][i])
			i++
			j--
		}
		if i == j { // 当 len(A[k]) When the length is odd, the number in the middle is processed
			A[k][i] = invert(A[k][i])
		}
	}
	return A
}

func flipAndInvertImageDebug(A [][]int) [][]int {
	for k := 0; k < len(A); k++ { // row
		i, j := 0, len(A[k])-1 // column
		for i < j {
			//A[k][i],A[k][j] = A[k][j],A[k][i]
			i++
			j--
		}
		if i == j { // 当 len(A[k]) When the length is odd, the number in the middle is processed
			A[k][i] = (A[k][i])
		}
	}
	return A
}

func invert(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}
