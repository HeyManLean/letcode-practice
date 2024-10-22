package main

import "code/utils"

/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。

更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。

1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums 已按 非严格递增 排列

https://leetcode.cn/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150
*/

func removeDuplicates(nums []int) int {
	/*
		使用快慢指针，慢指针对应结果位置，快指针做对比过滤元素

		原有顺序保持，一般需要双指针！
	*/
	var (
		k int = 0
		m int = 1
	)

	for m < len(nums) {
		if nums[k] == nums[m] {
			m += 1
		} else {
			if m > k+1 {
				nums[k+1] = nums[m]
			}
			k += 1
			m += 1
		}
	}
	return k + 1
}

func main() {
	utils.ArrayCheck(
		removeDuplicates,
		utils.ArrayItem{Nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, Expected: []int{0, 1, 2, 3, 4}},
	)
}
