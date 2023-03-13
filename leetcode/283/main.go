package main

func MoveZeroes(nums []int) {
	var fast, slow = 0, 0
	for fast < len(nums) {
		switch {
		case fast == slow:
			fast++
		case nums[slow] != 0:
			slow++
		case nums[fast] == 0:
			fast++
		default:
			nums[fast], nums[slow] = nums[slow], nums[fast]
			fast++
			slow++
		}
	}
}
