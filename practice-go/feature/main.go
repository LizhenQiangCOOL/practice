package feature

// 输入： [1,2,3,3,4,5,6,7]   6
// 输出 index [0,5]  // index [2,3]

// Solution3 不排序
// [1,2,3,3,4,5,6,7]  6
// [1,2] [1,3] [1,3] .... [1,7]
// [2,3] [2,3]...[2,7]
func Solution3(Arr []int, A int) (int, int) {
	m := make(map[int]int)
	for i := 0; i < len(Arr); i++ {
		m[Arr[i]] = i
	}
	for k, v := range m {
		searchValue := A - k
		if result, ok := m[searchValue]; ok {
			return v, result
		}
	}
	return -1, -1
}
