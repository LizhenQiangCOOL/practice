package main

import "sort"

var result [][]int
var track []int
var trackSum int

func combinationSum2(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	track = make([]int, 0)
	trackSum = 0
	//需要排序
	sort.Ints(candidates)
	backtrack(candidates, target, 0)
	return result
}

func backtrack(candidates []int, target int, start int) {
	// 满足条件
	if trackSum == target {
		//说明相加已经等于target
		tmp := make([]int, len(track))
		copy(tmp, track)
		result = append(result, tmp)
	}
	//回溯, 遍历所有可能
	for i := start; i < len(candidates); i++ {
		//剪枝，不包含重复组合
		if i-start > 0 && candidates[i] == candidates[i-1] {
			//说明前面元素和当前元素相同，而且前面元素已经处理了
			// 当前元素不用进入，避免重复项
			continue
		}
		// 剪枝, 后面的就不用做选择
		if trackSum > target {
			return
		}

		//做选择
		track = append(track, candidates[i])
		trackSum += candidates[i]
		backtrack(candidates, target, i+1)
		//回撤选择
		track = track[:len(track)-1]
		trackSum -= candidates[i]
	}

}
