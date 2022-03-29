package main

// Longest Common Prefix

func LongestCommonPrefix(arr []string) string {
	if len(arr) < 2 {
		return ""
	}
	var prefix string
	for i, str := range arr {
		if i == 0 {
			prefix = str
		} else {
			prefix = twoStringLongestCommonPrefix(prefix, str)
		}
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func twoStringLongestCommonPrefix(str1, str2 string) string {
	index := 0
	for index < len(str1) && index < len(str2) && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}
