package main

var result [][]int
var track []int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	track = make([]int, 0)
	backtrack(nums, 0)
	return result
}

func backtrack(nums []int, start int) {
	//满足条件，添加结果
	tmp := make([]int, len(track))
	copy(tmp, track)
	result = append(result, tmp)

	//回溯
	for i := start; i < len(nums); i++ {
		//做选择
		track = append(track, nums[i])
		// 遍历子树
		backtrack(nums, i+1)
		//撤回选择
		track = track[:len(track)-1]
	}
}
