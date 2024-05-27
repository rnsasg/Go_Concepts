package main

import (
	"container/list"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack *list.List
}

func Constructor(root *TreeNode) BSTIterator {
	stack := list.New()
	bst := BSTIterator{
		stack: stack,
	}
	bst.pushAll(root)
	return bst
}

func (this *BSTIterator) Next() int {
	temp := this.stack.Front()
	this.stack.Remove(temp)
	this.pushAll(temp.Value.(*TreeNode).Right)
	return temp.Value.(*TreeNode).Val
}

func (this *BSTIterator) HasNext() bool {
	if this.stack.Len() == 0 {
		return false
	}
	return true
}

func (this *BSTIterator) pushAll(root *TreeNode) {
	for root != nil {
		this.stack.PushFront(root)
		root = root.Left
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
