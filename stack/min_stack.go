package stack

// Notice: Keep min variable in each element for current scope, new elemenet is new scope
// Notice: Condition and variables check

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	min := val
	if len(this.stack) > 0 {
		min = Min(this.minStack[len(this.minStack)-1], val)
	}

	this.stack = append(this.stack, val)
	this.minStack = append(this.minStack, min)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
