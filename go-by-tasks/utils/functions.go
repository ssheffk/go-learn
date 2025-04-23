package utils

import "fmt"

// Higher Order Function
var double = func(n int) int { return n * 2 }

func apply(f func(int) int, x int) int {
	return f(x)
}

// Variadic Function (accept many args)
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func counter() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func simplySwap(a, b string) (string, string) {
	return b, a
}

func FunctionsDemo() {
	// closure counter
	next := counter()
	val1, val2 := simplySwap("front", "back")
	fmt.Println("HOC: ", apply(double, 5))
	fmt.Println("VF: ", sum(1, 2, 3, 4))
	fmt.Println("Recursion: Factorial(5) ", factorial(5))
	fmt.Println("Closure: ", next())
	fmt.Println("Closure: ", next())
	fmt.Println("Closure: ", next())
	fmt.Println("SimplySwap: ", val1, val2)
}
