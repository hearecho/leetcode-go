package palindrome_num

import "strconv"

//简单思路
//转换为字符串
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	for i:=0;i<l;i++ {
		if s[i:i+1] != s[l-i-1:l-i] {
			return false
		}
	}
	return true
}

//进阶思路
//不转换字符串,逆置后面数字和前面数字比较，如果相等或者是与长的除以相等则返回true
//需要在进行改进
func optimaize_isPalindrome(x int) bool  {
	if x< 0 {
		return false
	}
	if x < 10 {
		return true
	}
	temp := 0
	for x != 0{
		temp = temp *10 + x%10
		x = x/10
		if temp == x || (x/10 != 0 && temp == x/10) {
			return true
		}
	}
	return false
}
