package dailyProblem

//https://leetcode-cn.com/problems/reorganize-string/

func reorganizeString(s string) string {
	res := ""
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
	}
	//之后重新排序，每次从最多的两个中选择两个进行排序，如果最后
	return res
}
