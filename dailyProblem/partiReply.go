package dailyProblem

import "fmt"

//分割回文字符串，使得每个子串都是回文串。返回所有分割方案

// 回溯  每次截取子串的长度从 1 到当前剩余子串的长度，进行循环
//
func partReply(s string) [][]string {
	res := make([][]string, 0)
	bk(s, &res, make([]string, 0))
	return res
}

func bk(remainStr string, res *[][]string, curArr []string) {
	if len(remainStr) == 0 {
		temp := curArr
		curArr = curArr[0:0]
		fmt.Println(temp)
		*res = append(*res, temp)
		return
	}
	for i := 1; i <= len(remainStr); i++ {
		temp := remainStr[0:i]
		if isReply(temp) {
			//curArr = append(curArr, temp)
			//fmt.Println("回溯前:", curArr)
			bk(remainStr[i:], res, append(curArr, temp))
			//还原
			//curArr = curArr[:len(curArr)-1]
			//fmt.Println("回溯后:", curArr)
		}
	}
}

//判断是否是回文
func isReply(s string) bool {
	if len(s) == 1 {
		return true
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
