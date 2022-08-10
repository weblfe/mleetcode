package code

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var (
		caseList = []struct {
			target int
			list   []int
			exp    []int
		}{
			{
				target: 9,
				list:   []int{2, 7, 11, 15},
				exp:    []int{0, 1},
			},
			{
				target: 6,
				list:   []int{3, 2, 4},
				exp:    []int{1, 2},
			},
			{
				target: 6,
				list:   []int{3, 3},
				exp:    []int{0, 1},
			},
		}
	)
	for i, v := range caseList {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			rs := TwoSum(v.list, v.target)
			if !eq(rs, v.exp) {
				t.Errorf("two sum %v ,target: %d , exp: %v, but give : %v", v.list, v.target, v.exp, rs)
			}
		})
	}
}

func eq(arr []int, arr2 []int) bool {
	if len(arr) != len(arr2) {
		return false
	}
	for i, v := range arr {
		if v != arr2[i] {
			return false
		}
	}
	return true
}
