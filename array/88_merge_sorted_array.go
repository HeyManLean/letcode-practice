package main

import "fmt"

/* 88 合并两个有序数组

给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

示例 1：
输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]

https://leetcode.cn/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	/* 从最后一个位置往前，取两个数组最大值（数组最右边的值）填入目标数组最右侧，直到其中一个数组元素为0

	 */
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m -= 1
		} else {
			nums1[m+n-1] = nums2[n-1]
			n -= 1
		}
	}
	for n > 0 {
		nums1[n-1] = nums2[n-1]
		n -= 1
	}
}

func main() {
	var (
		nums1  = []int{1, 2, 3, 0, 0, 0}
		m      = 3
		nums2  = []int{2, 5, 6}
		n      = 3
		result = []int{1, 2, 2, 3, 5, 6}
	)
	merge(nums1, m, nums2, n)

	for i, num := range result {
		if nums1[i] != num {
			panic("not equal")
		}
	}
	fmt.Println("OK")
}
