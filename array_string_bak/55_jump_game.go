package main

import "code/test_utils"

/*
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。


输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

https://leetcode.cn/problems/jump-game/description/?envType=study-plan-v2&envId=top-interview-150
*/

func canJump(nums []int) bool {
	/*
		贪心算法

		条件判断：选择当前能够触及的位置，且结合位置跳跃长度之和最大的位置
	*/
	rightMost := 0

	for i := 0; i < len(nums); i++ {
		if i <= rightMost {
			if i+nums[i] > rightMost {
				rightMost = i + nums[i]
			}
			if rightMost >= len(nums)-1 {
				return true
			}
		}

	}
	return false
}

func testCanJump(nums []int) int {
	if canJump(nums) {
		return 1
	}
	return 0
}

func main() {
	test_utils.ArrayCheckResult(
		testCanJump,
		test_utils.ArrayItem{Nums: []int{2, 3, 1, 1, 4}, Result: 1},
		test_utils.ArrayItem{Nums: []int{3, 2, 1, 0, 4}, Result: 0},
	)
}
