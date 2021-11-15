package StrToInt

import (
	"math"
	"strings"
)

/**
超出范围的话返回最大值或者是最小值
数值范围为 [−231,  231 − 1]。如果数值超过这个范围，请返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。
*/

func strToInt(str string) int {
	//Flag 符号只能有一个或者没有 大于2则直接返回
	// ! 还有小数 只保留整数部分
	res := 0
	// 如果有空格先去除空格
	str = strings.TrimSpace(str)
	//然后开始遍历字符串
	flagNum := 0
	flag := false
	haveNum := false
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			break
		}
		// 碰到字符直接返回数字即可 或者符号数目大于1
		if str[i] >= 'a' && str[i] <= 'z' || flagNum > 1 {
			// 还有正负号
			if flag {
				return -res
			}
			return res
		}
		if str[i] == '+' || str[i] == '-' {
			if haveNum {
				break
			}
			if str[i] == '-' {
				flag = true
			}
			flagNum++
			continue
		}
		if str[i] < '0' || str[i] > '9' {
			break
		}
		haveNum = true
		temp := int(str[i] - '0')
		res = res*10 + temp
		if flag && -res <= math.MinInt32 {
			return math.MinInt32
		} else if !flag && res >= math.MaxInt32 {
			return math.MaxInt32
		}
	}
	if flag {
		return -res
	}
	return res
}
