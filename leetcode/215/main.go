package main

func patition(A []int, p, r int) int {
	x := A[r]
	slow := p - 1
	for fast := p; fast < r; fast++ {
		if A[fast] < x {
			slow++
			A[fast], A[slow] = A[slow], A[fast]
		}
	}
	A[slow+1], A[r] = A[r], A[slow+1]
	return slow + 1
}

func findKthLargest(nums []int, k int) int {
	x, left, right := 0, 0, len(nums)-1
	for {
		x = patition(nums, left, right)
		if x == len(nums)-k {
			break
		} else if x < len(nums)-k {
			//说明结果在右边
			left = x + 1
		} else {
			//说明结果在左边
			right = x - 1
		}
	}
	return nums[x]
}
