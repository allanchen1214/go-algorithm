package main

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

func Fibonacci_optimize(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	pre := 1
	ret := 1
	var tmp int
	for i := 2; i < n; i++ {
		tmp = ret
		ret = ret + pre
		pre = tmp
	}
	return ret
}
