package code

import "sort"

/// 本题的思路
/// 顺向： 1. 遍历求和
/// 逆向： 1. 遍历求差

/// 优化思路推导: 1. 减少数组遍历次数，最好只对每个值 遍历一次 -> 数据遍历后就缓存起来
///             2. 遍历次数减少 , 使用缓存结构 -> 空间换时间
///             3. 要便于查询 值 与  数组下标索引关系的数据结构 map 缓存，本题查询值为 数组值，结果期望为 索引下标
///             4. 所以 推导出 使用 map[Value]Index 存储关系 kv map
/// 						5. 小学 2项 加减法， 已知 和 与 加数，求另一个加数

func TwoSum(nums []int, target int) []int {
	var bucket = make(map[int]int)
	for k, v := range nums {
		if i, ok := bucket[target-v]; ok {
			return []int{i, k}
		}
		bucket[v] = k
	}
	return []int{}
}

func TwoSum2(nums []int, target int) []int {
	var (
		count = len(nums)
	)
	// 有序数组数组，sum
	sort.Ints(nums)
	for i, j := 0, count-1; i < j; {
		sum := nums[i] + nums[j]
		if sum == target {
			return []int{i, j}
		}
		if sum > target {
			j--
		} else {
			i++
		}
	}
	return []int{}
}
