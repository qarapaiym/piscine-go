package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	it := l.Head
	count := 0
	for it != nil {
		count++
		it = it.Next
	}
	return count
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	n.Next = l.Head
	l.Head = n
}
