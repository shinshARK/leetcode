type MinStack struct {
    Stack []int
    MinStack []int
}


func Constructor() MinStack {
    return MinStack{make([]int, 0, 10), make([]int, 0, 10)}
}


func (this *MinStack) Push(val int)  {
    this.Stack = append(this.Stack, val)
    if len(this.MinStack) == 0 || val <= this.MinStack[len(this.MinStack) - 1] {
        this.MinStack = append(this.MinStack, val)
    }
}


func (this *MinStack) Pop()  {
    if this.Stack[len(this.Stack) - 1] == this.MinStack[len(this.MinStack) - 1] {
        this.MinStack = this.MinStack[:len(this.MinStack) - 1]
    }
    this.Stack = this.Stack[:len(this.Stack) - 1]
}


func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack) - 1]
}


func (this *MinStack) GetMin() int {
    return this.MinStack[len(this.MinStack) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */