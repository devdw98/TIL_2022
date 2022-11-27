package src

func Fibonacci1(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	return Fibonacci1(n-1) + Fibonacci1(n-2)
}

func Fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}

	one := 1
	two := 0
	result := 0
	for i := 2; i <= n; i++ {
		result = one + two
		two = one
		one = result
	}
	return result
}
