package leetcode

import (
	"github.com/techlandspace/algorithm-go/internal/leetcode"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &leetcode.ListNode{Val: 2, Next: &leetcode.ListNode{Val: 4, Next: &leetcode.ListNode{Val: 3}}}
	l2 := &leetcode.ListNode{Val: 5, Next: &leetcode.ListNode{Val: 6, Next: &leetcode.ListNode{Val: 4}}}
	l3 := leetcode.AddTwoNumbers(l1, l2)
	for l3 != nil {
		t.Log(l3.Val)
		l3 = l3.Next
	}
}
