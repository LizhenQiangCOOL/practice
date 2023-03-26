package main

var res [][]int
var track []int
var used []bool

func permute(nums []int) [][]int {
	//leetcode 多用例测试
	res = make([][]int, 0)
	track = make([]int, 0)
	used = make([]bool, len(nums))
	backtrack(nums)
	return res
}

func backtrack(nums []int) {
	if len(track) == len(nums) {
		//说明已经到结束状态
		// go 要拷贝下，防止子切片都引用同一个切片
		temp := make([]int, len(nums))
		copy(temp, track)
		res = append(res, temp)
		return
	}
	for i := range nums {
		if used[i] {
			// 已经走过的路径
			continue
		}
		//做选择
		track = append(track, nums[i])
		used[i] = true
		//回溯
		backtrack(nums)
		//撤销选择
		track = track[:len(track)-1]
		used[i] = false
	}
}
