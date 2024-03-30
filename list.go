package gulc


// reverseList 翻转链表
func ReverseList(head *ListNode) *ListNode {
	res := head
	tr := head.Next
	head.Next = nil
	for tr != nil {
		tmp := tr.Next
		tr.Next = res
		res = tr
		tr = tmp
	}
	return res
}