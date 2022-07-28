package twopointers

import "github.com/bizshuk/algo/util"

// [Variant]: [TwoPointers] move less value pointer fisrt
func MaxArea(heights []int) int {
	mostWater := 0
	lp, rp := 0, len(heights)-1

	for lp < rp {
		water := util.Min(heights[lp], heights[rp]) * (rp - lp)
		mostWater = util.Max(mostWater, water)

		if heights[lp] > heights[rp] {
			rp -= 1
		} else {
			lp += 1
		}

	}
	return mostWater
}
