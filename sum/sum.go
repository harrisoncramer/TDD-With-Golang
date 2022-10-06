package sum

import "reflect"

// Product multiplys together all of the inputs and returns the result
func Product(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	for _, number := range nums[1:] {
		res *= number
	}
	return res
}

// ProductAll multiplys together the values in multiple sets of slices
func ProductAll(numbersToSum ...[]int) int {
	var nonEmptyArrays [][]int
	for _, arr := range numbersToSum {
		if !reflect.DeepEqual(arr, []int{}) {
			nonEmptyArrays = append(nonEmptyArrays, arr)
		}
	}

	if len(nonEmptyArrays) == 0 {
		return 0
	}

	res := Product(nonEmptyArrays[0])

	if len(nonEmptyArrays) > 1 {
		for _, s := range numbersToSum[1:] {
			adder := Product(s)
			if adder != 0 {
				res *= Product(s)
			}
		}
	}

	return res
}
