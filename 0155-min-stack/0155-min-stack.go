type MinStack struct {
    slice []int
    min []int
}


func Constructor() MinStack {
    return MinStack{make([]int, 0), make([]int, 0)}
}


func (this *MinStack) Push(val int)  {
    this.slice = append(this.slice, val)
    if len(this.min) == 0 || val <= this.min[len(this.min) - 1] {
        this.min = append(this.min, val)
    }
}


func (this *MinStack) Pop()  {
    if this.slice[len(this.slice) - 1] == this.min[len(this.min) - 1] {
        this.min = this.min[:len(this.min) - 1]
    }
    this.slice = this.slice[:len(this.slice) - 1]
}


func (this *MinStack) Top() int {
    return this.slice[len(this.slice) - 1]
}


func (this *MinStack) GetMin() int {
    return this.min[len(this.min) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */