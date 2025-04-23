package utils

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")

		default:
			fmt.Println(i)
		}
	}
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Calculate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b != 0 {
			return a / b
		}
		fmt.Println("Cannot divide by zero")
		return 0
	default:
		fmt.Println("Invalid operator")
		return 0
	}
}

func ControlFlowDemo() {
	fmt.Println("FizzBuzz")
	FizzBuzz(3)
	fmt.Println("IsPrime: ", IsPrime(13))
	fmt.Println("Calculate", Calculate(5, 3, "-"))
}
