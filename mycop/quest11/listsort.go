package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	i := l
	for i != nil {
		if i.Next != nil && i.Data > i.Next.Data {
			i.Data, i.Next.Data = i.Next.Data, i.Data
			i = l
		} else {
			i = i.Next
		}
	}
	return l
}
