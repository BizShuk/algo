package bit

// Find single number, all are duplicate once except one

// Notice: XOR(^) true-false
// XOR is
// true , true , 0
// true , false, 1
// false, true , 1
// false, false, 0

// If same bit => reset to 0
func xor(nums []int) int {
	xor_result := 0

	for _, num := range nums {
		xor_result ^= num
	}
	return xor_result
}
