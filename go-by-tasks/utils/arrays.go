package utils

import "fmt"

func RotateLeftArray(slice []int, n int) []int {
	n = n % len(slice)
	return append(slice[n:], slice[:n]...)
}

func RemoveDuplicates(nums []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			result = append(result, num)
		}
	}
	return result
}

func FindMinMax(nums []int) (int, int) {
	if len(nums) == 0 {
		panic("empty slice")
	}

	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func Flatten(matrix [][]int) []int {
	flat := []int{}

	for _, row := range matrix {
		flat = append(flat, row...)
	}
	return flat
}

func EfficientAppend() {
	s := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	fmt.Println(len(s), cap(s))
}

func DemoArrays() {
	fmt.Println("### Arrays")

	fmt.Println("RotateLeft:", RotateLeftArray([]int{1, 2, 3, 4, 5}, 2))     // [3 4 5 1 2]
	fmt.Println("RemoveDuplicates:", RemoveDuplicates([]int{1, 2, 2, 3, 1})) // [1 2 3]
	min, max := FindMinMax([]int{3, 1, 9, 2})
	fmt.Println("FindMinMax:", min, max) // 1 9
	matrix := [][]int{{1, 2}, {3, 4}, {5}}
	fmt.Println("Flatten:", Flatten(matrix)) // [1 2 3 4 5]

	fmt.Println("EfficientAppend:")
	EfficientAppend()
}
