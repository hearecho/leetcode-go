# [6. Z 字形变换](https://leetcode-cn.com/problems/zigzag-conversion/)

将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

> L   C   I   R
> E T O E S I I G
> E   D   H   N

之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

### 解题思路

```go
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
```

