package stack

import "sort"

// Notice: How to sort with other same-indexed array
// Notice: with accending position, the remaining time can be used for whetehr car bump into another

type Pair struct {
	remain   float64
	position int
}

func carFleet(target int, position []int, speed []int) int {
	pairs := make([]Pair, len(position))
	for i := range position {
		pairs[i] = Pair{float64(target-position[i]) / float64(speed[i]), position[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].position < pairs[j].position
	})

	stack := []Pair{}

	for _, v := range pairs { // Alternative: use prevTime to comapre might be simpler
		for len(stack) > 0 {
			if stack[len(stack)-1].remain <= v.remain {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}

		stack = append(stack, v)
	}

	return len(stack)
}
