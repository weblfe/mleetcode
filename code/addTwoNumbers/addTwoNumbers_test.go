package code

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	var (
		cases = []struct {
			node1 *ListNode
			node2 *ListNode
			node3 *ListNode
		}{
			{
				node1: NewNode(2, 4, 3),
				node2: NewNode(5, 6, 4),
				node3: NewNode(7, 0, 8),
			},
			{
				node1: NewNode(0),
				node2: NewNode(0),
				node3: NewNode(0),
			},
			{
				node1: NewNode(1),
				node2: NewNode(0),
				node3: NewNode(1),
			},
			{
				node1: NewNode(9, 9, 9, 9, 9, 9, 9),
				node2: NewNode(9, 9, 9, 9),
				node3: NewNode(8, 9, 9, 9, 0, 0, 0, 1),
			},
		}
	)

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d-%v+%v", i, c.node1, c.node2), func(t *testing.T) {
			tc := AddTwoNumbers(c.node1, c.node2)
			if !tc.Equal(c.node3) {
				t.Errorf("%v + %v ,expect node %v, but give %v", c.node1, c.node2, c.node3, tc)
			}
		})
	}
}
