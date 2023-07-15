package main

import "fmt"

type LinkNode struct {
	value int
	next *LinkNode
}

func NewLinkNode(value int) *LinkNode {
	return &LinkNode{
		value: value,
		next: nil,
	}
}

func (n *LinkNode) GetValue() int {
	return n.value
}

func (n *LinkNode) SetNext(next *LinkNode) {
	n.next = next
}

func (n *LinkNode) GetNext() *LinkNode {
	return n.next
}

func dump(head *LinkNode) {
	for {
		if head == nil {
			fmt.Printf("nil")
			break
		} else {
			fmt.Printf("%v=>", head.GetValue())
			head = head.GetNext()
		}
	}
	return
}

func main(){
	var node [5]*LinkNode
	for i := 0; i < len(node); i++ {
		node[i] = NewLinkNode(i)
	}
	for i := 0; i < len(node) - 1; i++ {
		node[i].SetNext(node[i + 1])
	}

	dump(node[0])

}