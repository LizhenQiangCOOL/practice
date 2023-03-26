package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// 非递归
func fibY(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	if n > 1 {
		for i := 2; i <= n; i++ {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}
