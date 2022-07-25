package stack

// Notice: Monotonic Stack(know index) ~= Priority Queue(should contain index) ~= Heap

func dailyTemperatures(temperatures []int) []int {
	stack := []int{}
	result := make([]int, len(temperatures))

	for i := range temperatures {

		for len(stack) > 0 { // pop all less than current temp
			j := stack[len(stack)-1]
			if temperatures[i] > temperatures[j] {
				result[j] = i - j
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		stack = append(stack, i)
	}

	return result
}
