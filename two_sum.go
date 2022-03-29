package main

func TwoSum(arr []int, sum int) (int, int) {
	m := make(map[int]int)
	for i, item := range arr {
		if val, ok := m[sum-item]; ok {
			return val, i
		}
		m[item] = i
	}
	return -1, -1
}
