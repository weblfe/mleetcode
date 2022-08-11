package reverseList_test

import (
	"fmt"
	. "mleetcode/code/reverseList"
	"testing"
)

func TestReverseList(t *testing.T) {
	var cases = []struct {
		src    *ListNode
		expect *ListNode
	}{

		{
			src:    NewListNode(1),
			expect: NewListNode(1),
		},
		{
			src:    NewListNode(1, 2, 5),
			expect: NewListNode(5, 2, 1),
		},
		{
			src:    NewListNode(1, 2, 3, 4, 5),
			expect: NewListNode(5, 4, 3, 2, 1),
		},
		{
			src:    NewListNode(),
			expect: NewListNode(),
		},
	}
	for _, v := range cases {
		// case 隔离 - 并非并发测试
		t.Run(fmt.Sprintf("%v-%v", v.src.String(), v.expect.String()), func(t *testing.T) {
			re := ReverseList(v.src)
			if re != nil && !re.Equal(v.expect) {
				t.Errorf("expect : %v", v.expect)
			}
		})
	}
}
