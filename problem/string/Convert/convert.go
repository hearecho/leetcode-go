package Convert

func convert(s string, numRows int) string {
	//创建 numRows行，或者是len(s)行字符串，最后将所有字符串相加在一起
	//由于可能字符串的长度小于 numRows的长度
	//所以 min(len(s),numRows)
	strs := make([]string, min(numRows, len(s)))
	//此时的行号
	rowIndex := 0
	//方向 开始向下
	direct := true
	for i := 0; i < len(s); i++ {
		//要决定什么时候转弯 由向下转到向上
		strs[rowIndex] += string(s[i])
		if direct {
			rowIndex++
		} else {
			rowIndex--
		}
		//当rowIndex等于numRows-1或者等于0的时候 转向
		if rowIndex == 0 || rowIndex == numRows-1 {
			direct = !direct
		}
	}
	res := ""
	for _, str := range strs {
		res += str
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return a
	}
	return b
}
