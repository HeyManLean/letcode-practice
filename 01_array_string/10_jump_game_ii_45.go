package main

import "code/test_utils"

/*
45. 跳跃游戏 II

给定一个长度为 `n` 的 **0 索引** 整数数组 `nums`。初始位置为 `nums[0]`。

每个元素 `nums[i]` 表示从索引 `i` 向前跳转的最大长度。换句话说，如果你在 `nums[i]` 处，你可以跳转到任意 `nums[i + j]` 处:

- `0 <= j <= nums[i]`
- `i + j < n`

返回到达 `nums[n - 1]` 的最小跳跃次数。生成的测试用例可以到达 `nums[n - 1]`。

**示例 1:**

```
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
     从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。

```

**示例 2:**

```
输入: nums = [2,3,0,1,4]
输出: 2

```

**提示:**

- `1 <= nums.length <= 104`
- `0 <= nums[i] <= 1000`
- 题目保证可以到达 `nums[n-1]`

https://leetcode.cn/problems/jump-game-ii/description
*/

func jump(nums []int) int {
	/*
		计算当前下标能到达的最大距离，当走到当前下标最大距离，则加一跳
	*/
	rightMost := 0
	end := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		rightMost = max(rightMost, i+nums[i])
		if i == end {
			end = rightMost
			steps++
		}
	}
	return steps
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func jumpDP(nums []int) int {
	/*
		动态规划

		假设当前位置最小跳跃数已经计算，来计算当前位置所及的位置最小跳跃数
	*/
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 0; i < len(nums)-1; i++ {
		// 不可达
		if i != 0 && dp[i] == 0 {
			continue
		}
		for j := 1; j <= nums[i]; j++ {
			if i+j > len(nums)-1 {
				break
			}
			if dp[i+j] == 0 || dp[i]+1 < dp[i+j] {
				dp[i+j] = dp[i] + 1
			}
		}
	}
	return dp[len(nums)-1]
}

func main() {
	test_utils.ArrayCheckResult(
		jump,
		test_utils.ArrayItem{Nums: []int{2, 3, 1, 1, 4}, Result: 2},
		test_utils.ArrayItem{Nums: []int{2, 3, 0, 1, 4}, Result: 2},
		test_utils.ArrayItem{Nums: []int{2, 1}, Result: 1},
	)
}
