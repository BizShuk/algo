package bit

// Notice:
// >> , shift 1 bit right
// << , shift 1 bit left
// >>> , No this operator
// <<< , No this operator
// & , n &= (n-1) => one time '1' bit count

func BitShift(num int) int /* Counter of '1' bit */ {
	counter := 0

	for num != 0 {
		num &= (num - 1)
		counter += 1
	}
	return counter
}
