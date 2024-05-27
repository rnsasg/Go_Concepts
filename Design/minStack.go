package main

import "container/list"

func main() {

}

type MinStack struct {
	stack    *list.List
	minStack *list.List
}

func Constructor() MinStack {
	stack := list.New()
	minStack := list.New()
	return MinStack{
		stack:    stack,
		minStack: minStack,
	}
}

func (this *MinStack) Push(val int) {

	this.stack.PushFront(val)
	// Push to min element on top of stack
	if this.minStack.Len() == 0 {
		this.minStack.PushFront(val)
	} else {
		item := this.minStack.Front().Value.(int)
		if val < item {
			this.minStack.PushFront(val)
		} else {
			this.minStack.PushFront(item)
		}
	}
}

func (this *MinStack) Pop() {
	if this.stack.Len() != 0 {
		this.stack.Remove(this.stack.Front())
	}
	if this.minStack.Len() != 0 {
		this.minStack.Remove(this.minStack.Front())
	}
}

func (this *MinStack) Top() int {
	if this.stack.Len() != 0 {
		return this.stack.Front().Value.(int)
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.minStack.Len() != 0 {
		return this.minStack.Front().Value.(int)
	}
	return -1
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
