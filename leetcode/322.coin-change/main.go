package main

import "math"

var memo []int

func coinChange(coins []int, amount int) int {
	memo = make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		memo[i] = -666
	}
	return dp(coins, amount)
}

func dp(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if memo[amount] != -666 {
		return memo[amount]
	}
	res := math.MaxInt
	for _, v := range coins {
		subProblem := dp(coins, amount-v)
		if subProblem == -1 {
			continue
		}
		res = min(res, subProblem+1)
	}

	if res == math.MaxInt {
		memo[amount] = -1
		return -1
	}
	memo[amount] = res
	return res
}

func dpNew(coins []int, amount int) int {
	dpSlice := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dpSlice[i] = amount + 1
	}
	// base case
	dpSlice[0] = 0
	// 遍历所有状态的所有取值
	for i := 0; i < len(dpSlice); i++ {
		// 遍历在求所有选择的最小值
		for _, coin := range coins {
			//子问题无解
			if i-coin < 0 {
				continue
			}
			dpSlice[i] = min(dpSlice[i], 1+dpSlice[i-coin])
		}
	}
	if dpSlice[amount] == amount+1 {
		return -1
	}
	return dpSlice[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
