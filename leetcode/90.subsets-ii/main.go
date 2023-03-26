package main

import "sort"

var result [][]int
var track []int

func subsetsWithDup(nums []int) [][]int {
	result = make([][]int, 0)
	track = make([]int, 0)
	sort.Ints(nums)
	backtrack(nums, 0)
	return result
}

// start 从第几个开始选
func backtrack(nums []int, start int) {
	// 满足条件加入
	tmp := make([]int, len(track))
	copy(tmp, track)
	result = append(result, tmp)
	//回溯
	for i := start; i < len(nums); i++ {
		// 截肢
		if i > start && nums[i] == nums[i-1] {
			//不能和前面的相等
			continue
		}
		//做选择
		track = append(track, nums[i])
		backtrack(nums, i+1)
		//回撤
		track = track[:len(track)-1]
	}
}
