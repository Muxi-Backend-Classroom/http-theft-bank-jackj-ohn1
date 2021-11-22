package main

import "fmt"

func sum(num int) int {
	var sum = 1
	for i := 1; i <= num; i++ {
		sum = sum * i
	}
	return sum
}

func swap(a *int, b *int) {
	var tmp = *a
	*a = *b
	*b = tmp
}

var count = 0

func do(nums []int, m int, max int, re [][]int) [][]int {
	if m == max-1 {
		var tmp = make([]int, max)
		for i := 0; i < max; i++ {
			tmp[i] = nums[i]
		}
		re[count] = tmp
		count++
	} else {
		for i := m; i < max; i++ {
			swap(&nums[i], &nums[m])
			do(nums, m+1, max, re)
			swap(&nums[i], &nums[m])
		}
	}
	return re
}

func permute(nums []int) [][]int {
	tmp := make([][]int, sum(len(nums)))
	re := do(nums, 0, len(nums), tmp)
	return re
}

func main() {
	var num int
	fmt.Scanf("%d", &num)
	testSlice := make([]int, num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &testSlice[i])
	}
	res := permute(testSlice)
	fmt.Println(res)
}
