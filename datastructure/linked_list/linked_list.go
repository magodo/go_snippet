package linked_list

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func (l *ListNode) ToSlice() []interface{} {
	out := []interface{}{}
	for l != nil {
		out = append(out, l.Val)
		l = l.Next
	}
	return out
}

func FromSlice(slice []interface{}) *ListNode {
	nodeBeforeHead := &ListNode{}
	l := nodeBeforeHead
	for _, i := range slice {
		l.Next = &ListNode{Val: i}
		l = l.Next
	}
	return nodeBeforeHead.Next
}
