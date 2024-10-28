package main

import (
	"code/test_utils"
	"sort"
)

/*
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。
元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。

假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
返回 k。

0 <= nums.length <= 100
0 <= nums[i] <= 50
0 <= val <= 100

https://leetcode.cn/problems/remove-element/?envType=study-plan-v2&envId=top-interview-150
*/

func removeElement(nums []int, val int) int {
	/*
		左右各一个指针，左边元素遍历，如果等于 val，则跟右边的元素交换，直到两指针位置一样
		k 代表有多少个元素 != val
	*/
	if len(nums) == 0 {
		return 0
	}
	var (
		m int = 0
		k int = len(nums)
	)
	for m <= k-1 {
		if nums[k-1] == val {
			k -= 1
			continue
		}
		if nums[m] == val {
			nums[k-1], nums[m] = nums[m], nums[k-1]
			k -= 1
			m += 1
		} else {
			m += 1
		}
	}
	sort.Ints(nums[:k]) // 排序
	return k
}

func main() {
	test_utils.ArrayCheckWithVal(
		removeElement,
		test_utils.ArrayItem{Nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, Val: 2, Expected: []int{0, 0, 1, 3, 4}},
		test_utils.ArrayItem{Nums: []int{1}, Val: 1, Expected: []int{}},
	)
}
