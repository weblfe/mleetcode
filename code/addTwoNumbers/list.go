package code

import (
	"bytes"
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(el ...int) *ListNode {
	var (
		h    = &ListNode{Val: 0}
		iter = h
	)
	for _, v := range el {
		iter.Next = &ListNode{Val: v}
		iter = iter.Next
	}
	return h.Next
}

func (node *ListNode) Equal(n *ListNode) bool {
	if node == n {
		return true
	}
	var iter = node
	for iter != nil && n != nil {
		if iter.Val != n.Val {
			return false
		}
		iter = iter.Next
		n = n.Next
	}
	if iter == nil && n == nil {
		return true
	}
	return false
}

func (node *ListNode) String() string {
	var (
		iter   = node
		values []string
		buffer = bytes.NewBufferString(`[`)
	)
	for iter != nil {
		values = append(values, fmt.Sprintf("%d", iter.Val))
		iter = iter.Next
	}
	buffer.WriteString(strings.Join(values, ","))
	buffer.WriteString("]")
	return buffer.String()
}
