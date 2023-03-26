package main

var result [][]int
var track []int

func combine(n int, k int) [][]int {
	result = make([][]int, 0)
	track = make([]int, 0)
	backtrace(n, k, 1)
	return result
}

func backtrace(n int, k int, start int) {
	// 第k层结束
	if k == len(track) {
		tmp := make([]int, k)
		copy(tmp, track)
		result = append(result, tmp)
		return
	}
	// 回溯 ，start从1开始
	for i := start; i <= n; i++ {
		// 做选择
		track = append(track, i)
		backtrace(n, k, i+1)
		//回撤选择
		track = track[:len(track)-1]
	}
}
