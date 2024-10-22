package utils

import (
	"fmt"
)

type ArrayItem struct {
	Nums     []int // 作为参数传入
	Expected []int // 期望的结果数组
	Val      int   // 作为参数传入
	Result   int   // 处理后期望返回结果
}

func ArrayCheck(handler func([]int) int, items ...ArrayItem) {
	for _, item := range items {
		result := handler(item.Nums)
		if result != len(item.Expected) {
			panic(fmt.Sprintf("len not correct. len=%d, expected=%d", result, len(item.Expected)))
		}
		for i := 0; i < len(item.Expected); i++ {
			if item.Nums[i] != item.Expected[i] {
				panic(fmt.Sprintf("value not correct. value=%v, expected=%v", item.Nums[:result], item.Expected))
			}
		}
	}
	fmt.Println("OK.")
}

func ArrayCheckWithVal(handler func([]int, int) int, items ...ArrayItem) {
	for _, item := range items {
		result := handler(item.Nums, item.Val)
		if result != len(item.Expected) {
			panic(fmt.Sprintf("len not correct. len=%d, expected=%d", result, len(item.Expected)))
		}
		for i := 0; i < len(item.Expected); i++ {
			if item.Nums[i] != item.Expected[i] {
				panic(fmt.Sprintf("value not correct. value=%v, expected=%v", item.Nums[:result], item.Expected))
			}
		}
	}
	fmt.Println("OK.")
}

func ArrayCheckResult(handler func([]int) int, items ...ArrayItem) {
	for _, item := range items {
		result := handler(item.Nums)
		if result != item.Result {
			panic(fmt.Sprintf("result not correct. result=%d, expected=%d", result, item.Result))
		}
	}
	fmt.Println("OK.")
}
