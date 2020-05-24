package BigNumAdd

import "fmt"

// 反转链表，相加，注意进位
func (l *ListNode)MyAddTwoNumLink(l1 *ListNode) *ListNode {
	//head := new(ListNode)
	var head *ListNode
	head = &ListNode{}
	newList := new(ListNode)
	newList = head
	var newNode *ListNode
	l = l.myreverse()
	l1 = l1.myreverse()
	data := 0
	carry := 0
	//fmt.Println(l,l1)
	for l != nil && l1 != nil {
		fmt.Println(l.data, l1.data)
		res := l.data + l1.data + carry
		data = (res) % 10
		newNode = &ListNode{data: data}
		head.next = newNode
		carry = (res) / 10
		head = head.next
		l = l.next
		l1 = l1.next
	}

	for l != nil {
		head.next = l
		head = head.next
		l = l.next
	}

	for l1 != nil {
		head.next = l1
		head = head.next
		l1 = l1.next
	}

	return newList.next.myreverse()
}


func (l *ListNode)myreverse() *ListNode{
	tmp := new(ListNode)
	var pre *ListNode
	cur := l
	for cur != nil {
		tmp = cur.next  //拴住cur节点的next节点 防止丢失

		cur.next = pre  //cur节点的next重指向到pre节点，完成翻转

		pre = cur   //cur节点变成pre节点
		cur = tmp   //拴住的临时节点成了当前节点
	}
	return pre  //cur 节点 是 空节点
}
