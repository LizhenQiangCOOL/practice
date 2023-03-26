package main

var memo []int

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	memo = make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		res = max(lengthOfLISInner(nums, i), res)
	}
	return res
}

func lengthOfLISInner(nums []int, n int) int {
	//base case 无
	if memo[n] != 0 {
		return memo[n]
	}
	res := 1
	for i := 0; i <= n-1; i++ {
		if nums[i] >= nums[n] {
			continue
		}
		res = max(lengthOfLISInner(nums, i)+1, res)
	}
	memo[n] = res
	return res
}

// 非递归
func lengthOfLISInnerY(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	//遍历所有状态
	for i := 0; i < len(dp); i++ {
		//遍历所有选择
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 1
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
