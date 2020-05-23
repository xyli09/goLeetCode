package minStack

type MinStack struct {
	elems []int
	mins  []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(x int) {
	this.elems = append(this.elems, x)

	if len(this.mins) == 0 || this.GetMin() >= x {
		this.mins = append(this.mins, x)
	}
}

func (this *MinStack) Pop() {
	elem := this.Top()
	this.elems = this.elems[:len(this.elems)-1]

	if elem <= this.GetMin() {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.elems) == 0 {
		panic("empty stack")
	}

	elem := this.elems[len(this.elems)-1]
	return elem
}

func (this *MinStack) GetMin() int {
	if len(this.mins) == 0 {
		panic("empty stack")
	}

	elem := this.mins[len(this.mins)-1]
	return elem
}

