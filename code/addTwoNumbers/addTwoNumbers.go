package code

/// 解题思路
///  顺向 1. 逆序链表 遍历 加法 -> 简化 就是 正整数 珠算 低位 -> 高位 +，满十进一
///  解题功能点： 1. 逆序链表构造， 2. 有头和头链表 效率， 3. 算术 (取余，整除, 进位[个数加法])
///  注意事项 ： 返回结果也是 逆序链表

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var (
		carry   = 0
		head    = &ListNode{Val: 0}
		current = head
	)
	for carry != 0 || l1 != nil || l2 != nil {
		// 0,0,0 [shadow]
		var n1, n2, sum int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum = n1 + n2 + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	return head.Next
}
