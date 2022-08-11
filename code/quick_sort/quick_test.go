package quick_sort

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"testing"
)

type Cases[T constraints.Ordered] struct {
	src    []T
	expect []T
}

func TestQuickSortInt(t *testing.T) {
	var cases []Cases[int]
	cases = []Cases[int]{
		{
			src:    []int{1, 10, 9, 4, 8, 2, 3},
			expect: []int{1, 2, 3, 4, 8, 9, 10},
		},
		{
			src:    []int{},
			expect: []int{},
		},
		{
			src:    []int{-1, 0, 10, 4, 8},
			expect: []int{-1, 0, 4, 8, 10},
		},
		{
			src:    []int{1, 1, 1, 1, 1, 1, 1, 1},
			expect: []int{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			src:    []int{0},
			expect: []int{0},
		},
		{
			src:    []int{0, 1, 2, 3, 4, 5, 6},
			expect: []int{0, 1, 2, 3, 4, 5, 6},
		},
		{
			src:    []int{100, 1, 2, 3, 4, 5, 6},
			expect: []int{1, 2, 3, 4, 5, 6, 100},
		},
	}

	for _, v := range cases {
		t.Run(fmt.Sprintf("%v", v.src), func(t *testing.T) {
			src := QuickSort(v.src)
			t.Logf("%v\n", src)
			if fmt.Sprintf("%v", src) != fmt.Sprintf("%v", v.expect) {
				t.Errorf("expect %v", v.expect)
			}
		})
	}
}

func TestQuickSortFloat(t *testing.T) {
	var cases []Cases[float64]
	cases = []Cases[float64]{
		{
			src:    []float64{1, 10, 9, 4, 8, 2, 3},
			expect: []float64{1, 2, 3, 4, 8, 9, 10},
		},
		{
			src:    []float64{},
			expect: []float64{},
		},
		{
			src:    []float64{-1, 0, 10, 4, 8},
			expect: []float64{-1, 0, 4, 8, 10},
		},
		{
			src:    []float64{1, 1, 1, 1, 1, 1, 1, 1},
			expect: []float64{1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			src:    []float64{0},
			expect: []float64{0},
		},
	}

	for _, v := range cases {
		t.Run(fmt.Sprintf("%v", v.src), func(t *testing.T) {
			if fmt.Sprintf("%v", QuickSort(v.src)) != fmt.Sprintf("%v", v.expect) {
				t.Errorf("expect %v", v.expect)
			}
		})
	}
}
