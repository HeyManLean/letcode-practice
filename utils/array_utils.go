package utils

import (
	"fmt"
)

type ArrayItem struct {
	Nums     []int
	Expected []int
	Val      int
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
