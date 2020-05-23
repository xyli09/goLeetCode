package BigNumAdd

import "fmt"
import "strconv"
import "reflect"
type ListNode struct {
	data int
	next *ListNode
}
//转成字符串，数字 相加，再转成链表
func AddTwoNum(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Int,_ := strconv.Atoi(reverse(l1))
	l2Int,_ := strconv.Atoi(reverse(l2))
	l := l1Int + l2Int
	lStr := strconv.Itoa(l)
	fmt.Println(reflect.TypeOf(lStr))
	fmt.Println(lStr)
	var newList []*ListNode
	for _,v := range lStr {
		fmt.Println(string(v))
		n,_ := strconv.Atoi(string(v))
		node := &ListNode{data:n}
		newList = append(newList,node)
	}
	head := new(ListNode)
	tail := head
	for k,node := range newList {
		if k == len(newList)-1 {
			newList = append(newList,nil)
		}
		node.next = newList[k+1]
		tail.next = node
		tail = tail.next
	}
	return head
}

func reverse(l *ListNode) string {
	var s string
	for l != nil {
		s += strconv.Itoa(l.data)
		l = l.next
	}
	return s
}

// 反转链表，相加，注意进位
func (l *ListNode)AddTwoNumLink(l1 *ListNode) *ListNode {
	//head := new(ListNode)
	var head *ListNode
	head = &ListNode{}
	newList := new(ListNode)
	newList = head
	var newNode *ListNode
	l = l.reverseLink()
	l1 = l1.reverseLink()
	data := 0
	carry := 0
	//fmt.Println(l,l1)
	for l != nil || l1 != nil {
		var data1,data2 int
		if l != nil {
			data1 = l.data
		} else {
			data1 = 0
		}
		if l1 != nil {
			data2 = l1.data
		} else {
			data2 = 0
		}
		fmt.Println(data1,data2)
		data = (data1 + data2 + carry) % 10
		newNode = &ListNode{data:data}
		head.next = newNode
		carry = (data1 + data2 + carry) / 10
		head = head.next

		if l != nil {
			l = l.next
		}
		if l1 != nil {
			l1 = l1.next
		}
	}
	return newList.next.reverseLink()
}

func (l *ListNode)reverseLink() *ListNode {
	//pre := new(ListNode)
	var pre *ListNode
	next := new(ListNode)
	for l != nil {
		next = l.next
		//fmt.Println(l.next)
		l.next = pre
		pre = l
		l = next
	}
	l = pre
	return l
	//return pre
}

func (l *ListNode)readLink() {
	var result []int
	for l != nil {
		result = append(result,l.data)
		l = l.next
	}
	fmt.Println(result)
}

func (l *ListNode)myreverse(){
	tmp := new(ListNode)
	pre := l
	pre.next = nil
	cur := l.next
	for cur != nil&& cur.next != nil {
		if cur.next != nil{
			tmp = cur.next
			cur.next = pre
			pre = cur
			cur = tmp
		}
	}
	cur.next = pre
	l = pre
}

func testLinkNodeAdd()  {
	node7 := &ListNode{data:7,next:nil}
	node6 := &ListNode{data:6,next:node7}
	node5 := &ListNode{data:5,next:node6}

	node4 := &ListNode{data:4,next: node5}

	node3 := &ListNode{data:3,next:nil}
	//node3 := &ListNode{data:3,next:nil}
	node2 := &ListNode{data:2,next:node3}
	node1 := &ListNode{data:1,next: node2}


	//x := AddTwoNum(node1,node4)
	node1.readLink()
	//node1.reverseLink().readLink()
	node4.myreverse()
	node4.readLink()
	node1.AddTwoNumLink(node4).readLink()
	//fmt.Println(r)

}
