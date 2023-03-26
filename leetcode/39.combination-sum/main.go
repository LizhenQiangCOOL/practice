package main

import "sort"

var result [][]int
var track []int
var trackSum int

func combinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	track = make([]int, 0)
	trackSum = 0
	sort.Ints(candidates)
	backtrack(candidates, target, 0)
	return result
}

func backtrack(candidates []int, target int, start int) {
	//满足条件的
	if trackSum == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		result = append(result, tmp)
	}
	//回溯，选择所有可能,
	// 剪枝，选到后面了就不要选前面的了,从start开始
	for i := start; i < len(candidates); i++ {
		//剪枝, 已确定有序
		if trackSum+candidates[i] > target {
			//说明后面不用做选择，都已经超过target了
			break
		}
		//选择
		track = append(track, candidates[i])
		trackSum += candidates[i]
		backtrack(candidates, target, i)
		//回扯
		track = track[:len(track)-1]
		trackSum -= candidates[i]
	}
}
