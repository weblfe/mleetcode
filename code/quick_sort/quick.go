package quick_sort

import (
	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](v []T) []T {
	var count = len(v)
	if count <= 0 {
		return v
	}
	var low, high = 0, count - 1
	return quick(v, low, high)
}

func quick[T constraints.Ordered](v []T, low, high int) []T {
	if high > low {
		// 寻找基准点
		pivotIndex := partition(v, low, high)
		// 递归左边
		quick(v, low, pivotIndex-1)
		// 递归右边
		quick(v, pivotIndex+1, high)
	}
	return v
}

// 分区   [3][1], [1][3]
func partition[T constraints.Ordered](v []T, low, high int) int {
	pivot := v[low]
	// 游标是否已相交
	for low < high {

		// 假设最小数 小于 高位数，继续后移,往后比较
		if pivot <= v[high] {
			high--
		}

		v[low] = v[high] // [1][1], [3][3]
		// 假设最小 其实大于 他右边的数
		if low < high && pivot >= v[low] {
			low++
		}
		v[high] = v[low] // [1][1] , [3][3]
	}
	v[low] = pivot // [1][3], [1][3]
	return low     // 1
}

// [10] [1] [3] [2] [0] [8] [11]
// low=0 , high = 6

// 1. pivot = 10
// [10] [1] [3] [2] [0] [8] [11]
// 10 <= 11 ,low =0, high = 6 -1 -> high = 5
// [8] [1] [3] [2] [0] [8] [11]
// 10 >=8 , low = 0+1 -> low=1, high = 5
// [8] [1] [3] [2] [0] [1] [11]

// 10 <= 1 ,low =1, high = 5
// [8] [1] [3] [2] [0] [1] [11]

// 10 >=1 , low = 1+1 -> low=2, high = 5
// [8] [1] [3] [2] [0] [3] [11]

// 10 <= 3 , low = 2 , high = 5
// [8] [1] [3] [2] [0] [3] [11]

// 10 >=3 , low = 2+1 -> low=3, high = 5
// [8] [1] [3] [2] [0] [2] [11]

// 10 <= 2 , low = 3 , high = 5
// [8] [1] [3] [2] [0] [2] [11]

// 10 >= 2 , low = 3 +1 -> low =4 , high = 5
// [8] [1] [3] [2] [0] [0] [11]

// 10 <=0  , low = 4  , high = 5
// [8] [1] [3] [2] [0] [0] [11]

// 10 >=0  , low = 4 +1 -> low =5 , high = 5
// [8] [1] [3] [2] [0] [0] [11]

//  [8] [1] [3] [2] [0] [10] [11]

// -> low = 5
