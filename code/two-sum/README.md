# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## 题目

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```



## 题目大意

在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。

## 解题思路

这道题最优的做法时间复杂度是 O(n)。

顺序扫描数组，对每一个元素，在 map 中找能组合给定值的另一半数字，如果找到了，直接返回 2 个数字的下标即可。如果找不到，就把这个数字存入 map 中，等待扫到“另一半”数字的时候，再取出来返回结果。

## My 推导过程

```markdown
## 本题的思路
顺向： 1. 遍历求和
逆向： 1. 遍历求差

### 优化思路推导: 
1. 减少数组遍历次数，最好只对每个值 遍历一次 -> 数据遍历后就缓存起来
2. 遍历次数减少 , 使用缓存结构 -> 空间换时间
3. 要便于查询 值 与 数组下标索引关系的数据结构(kv) -> map 缓存，本题查询值为 数组值，结果期望为 索引下标
4. 所以 推导出 使用 map[Value]Index 存储关系 kv map
5. 小学 2项 加减法， 已知 和 与 加数，求另一个加数
```