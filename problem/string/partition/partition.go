package partition

/**
https://leetcode-cn.com/problems/palindrome-partitioning/
*/

func partition(s string) [][]string {
	// 一种是使用回溯方式 来进行分割字符串，之后判断字符串是否是回文串
	res := make([][]string, 0)
	var dfs func(left, right int, curStrs []string)
	dfs = func(left, right int, curStrs []string) {
		if right == len(s) {
			temp := make([]string, len(curStrs))
			copy(temp, curStrs)
			res = append(res, temp)
			return
		}
		for i := left + 1; i <= len(s); i++ {
			temp := s[left:i]
			if checkStr(temp) {
				dfs(i, i, append(curStrs, temp))
			}
		}
	}

	dfs(0, 0, make([]string, 0))
	return res
}

func checkStr(s string) bool {
	i, j := 0, len(s)-1
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
