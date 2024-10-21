package main

/*
给你一个有序数组 nums，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

更改数组 nums，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。

1 <= nums.length <= 3 * 104
-104 <= nums[i] <= 104
nums 已按 非严格递增 排列
*/
import (
	"code/utils"
)

func removeDuplicatesII(nums []int) int {
	/*
		双指针：slow 代表处理出的长度（结果），fast 代表已经检查过的长度
		对比 slow - 2 和 fast 是否一致，否则将 fast 的值赋值给 slow
	*/
	if len(nums) <= 2 {
		return len(nums)
	}
	slow := 2
	fast := 2

	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow += 1
		}
		fast += 1
	}
	return slow
}

func main() {
	utils.ArrayCheck(
		removeDuplicatesII,
		utils.ArrayItem{Nums: []int{1, 1, 1}, Expected: []int{1, 1}},
		utils.ArrayItem{Nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3}, Expected: []int{0, 0, 1, 1, 2, 3, 3}},
	)
}
