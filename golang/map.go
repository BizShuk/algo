package golang

// Notice: [Golang] add map key in loop will increase more loop cycle, but not for slice
func MapIncrementInLoop() {
	y := map[int]bool{0: true}
	i := 0
	for k := range y {
		i += 1
		y[k+1] = true
		if i > 10 {
			break
		}
	}
}
