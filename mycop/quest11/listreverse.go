package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
	//Prev *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	l_new := &List{}
	for i := l.Head; i != nil; i = i.Next {
		ListPushFront(l_new, i.Data)
	}
	*l = *l_new
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Next = n
	}
	l.Tail = n
}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	n.Next = l.Head
	l.Head = n
}
