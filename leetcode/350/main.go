package main

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	// 预先给定要装多少，内部直接算一个B，后续过程就不用扩容
	m := make(map[int]int, len(nums1))
	//预定内存，防止数组扩容拷贝
	result := make([]int, 0, len(nums1))
	for _, num := range nums1 {
		m[num]++
	}
	for _, num := range nums2 {
		if v, ok := m[num]; ok && v > 0 {
			result = append(result, num)
			m[num]--
		}
	}
	return result
}
