package array

// Notice: interesting flow of product

func productExceptSelf(nums []int) []int {
	product := 1
	result := make([]int, len(nums))

	for k, num := range nums {
		result[k] = product
		product *= num
	}
	// return result

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= product
		product *= nums[i]
	}
	return result
}
