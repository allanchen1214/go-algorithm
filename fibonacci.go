package main

// 这个方法性能超慢
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
	fn := 1
	var tmp int
	for i := 2; i < n; i++ {
		tmp = fn
		fn = pre + fn
		pre = tmp
	}
	return fn
}
