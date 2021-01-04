package main

type MinStack struct {
	minNums []int
	stack   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		minNums: []int{},
		stack:   []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if len(this.minNums) == 0 {
		this.minNums = append(this.minNums, x)
	} else {
		this.minNums = append(this.minNums, min(this.minNums[len(this.minNums)-1], x))
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minNums = this.minNums[:len(this.minNums)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minNums[len(this.minNums)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
