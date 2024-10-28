package main

import "code/test_utils"

/*
给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

示例 1:

输入: nums = [1,2,3,4,5,6,7], k = 3
输出: [5,6,7,1,2,3,4]
解释:
向右轮转 1 步: [7,1,2,3,4,5,6]
向右轮转 2 步: [6,7,1,2,3,4,5]
向右轮转 3 步: [5,6,7,1,2,3,4]

示例 2:
输入：nums = [-1,-100,3,99], k = 2
输出：[3,99,-1,-100]
解释:
向右轮转 1 步: [99,-1,-100,3]
向右轮转 2 步: [3,99,-1,-100]

https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-interview-150
*/

func rotate(nums []int, k int) {
	/*
		逆序遍历，使用延长数组，保存右移后数值
		再将延长数组赋值到原数组前面k位数
	*/
	extendNums := make([]int, k)
	for i := len(nums) - 1; i >= 0; i-- {
		if i+k >= len(nums) {
			extendNums[i+k-len(nums)] = nums[i]
		} else {
			nums[i+k] = nums[i]
		}
	}

	for i := 0; i < k; i++ {
		nums[i%len(nums)] = extendNums[i] // [-1], k=2
	}
}

func rotateWithResult(nums []int, k int) int {
	/*
		满足测试handler格式
	*/
	rotate(nums, k)
	return len(nums)
}

func main() {
	test_utils.ArrayCheckWithVal(
		rotateWithResult,
		test_utils.ArrayItem{Nums: []int{1, 2, 3, 4, 5, 6, 7}, Expected: []int{5, 6, 7, 1, 2, 3, 4}, Val: 3},
		test_utils.ArrayItem{Nums: []int{-1}, Expected: []int{-1}, Val: 2},
	)
}
