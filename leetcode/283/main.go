package main

func MoveZeroes(nums []int) {
	// nums 为 切片, 为引用类型
	var n int = len(nums)
	var fast, slow int = 0, 0
	for fast < n {
		if nums[fast] != 0 && nums[slow] == 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			fast++
			slow++
			continue
		}
		if fast == slow {
			fast++
			continue
		}
		if nums[fast] == 0 {
			fast++
		}
		if nums[slow] != 0 {
			slow++
		}
	}
}

func MoveZeroes_leetcode(nums []int) {
	var n int = len(nums)
	var fast, slow int = 0, 0
	for fast < n {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
