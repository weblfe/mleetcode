package reverseList

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(values ...int) *ListNode {
	var (
		head = &ListNode{}
		node = head
	)
	for _, v := range values {
		head.Next = &ListNode{
			Val: v,
		}
		head = head.Next
	}
	return node.Next
}

func (list *ListNode) String() string {
	var (
		iter   = list
		values []string
		str    = bytes.NewBufferString("[")
	)
	for iter != nil {
		values = append(values, fmt.Sprintf("%d", iter.Val))
		iter = iter.Next
	}
	str.WriteString(strings.Join(values, ","))
	str.WriteString("]")
	return str.String()
}

func (list *ListNode) Equal(li *ListNode) bool {
	if list == li {
		return true
	}
	if li == nil || list == nil {
		return false
	}
	return list.String() == li.String()
}

func (list *ListNode) HashCode() string {
	var hashUtil = md5.New()
	hashUtil.Write([]byte(list.String()))
	return fmt.Sprintf("%x", hashUtil.Sum(nil))
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 最后一个
	var last = ReverseList(head.Next)
	// 倒数第二个---指向前一个
	head.Next.Next = head
	// 断环
	head.Next = nil
	return last

}
