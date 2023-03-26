package main

import "sort"

var result [][]int
var track []int

// 防止遍历自己,记录哪些已经被遍历
var used []bool

func permuteUnique(nums []int) [][]int {
	result = make([][]int, 0)
	track = make([]int, 0)
	used = make([]bool, len(nums))
	sort.Ints(nums)
	backtrack(nums)
	return result
}

// start 是标准当前选项
func backtrack(nums []int) {
	// 满足条件
	if len(track) == len(nums) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		result = append(result, tmp)
	}
	// 回溯，遍历所有可能(全遍历)
	for i := 0; i < len(nums); i++ {
		if used[i] {
			// 当前已经被选择了,不用重复选择
			continue
		}
		//固定相同的元素在排列中的相对位置,
		// 即想要选择当前元素，保证前面元素已经被选
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			continue
		}
		//选择
		track = append(track, nums[i])
		used[i] = true
		backtrack(nums)
		track = track[:len(track)-1]
		used[i] = false
	}
}
