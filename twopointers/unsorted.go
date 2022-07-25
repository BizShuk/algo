package twopointers

// Notice: Levarage Max for maximum replacement
// Notice: two pointers tips: move pointer which has m/l
// if max < xxx { max = xxx }  => max = Max(max, xxx)

func maxArea(heights []int) int {
	mostWater := 0
	lp, rp := 0, len(heights)-1

	for lp < rp {
		water := Min(heights[lp], heights[rp]) * (rp - lp)
		mostWater = Max(mostWater, water)

		if heights[lp] > heights[rp] {
			rp -= 1
		} else {
			lp += 1
		}

	}
	return mostWater
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
