package easy

var fibs = make([]int,31)

func fib(N int) int {
	if N==1||N==0 {
		return N
	}

	res := fibs[N]
	if res == 0 {
		res = fib(N-1) + fib(N-2)
		fibs[N] = res
	}
	return res
}