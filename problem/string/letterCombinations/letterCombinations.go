package letterCombinations

var dict map[string]string = map[string]string{
	"2":"abc",
	"3":"def",
	"4":"ghi",
	"5":"jkl",
	"6":"mno",
	"7":"pqrs",
	"8":"tuv",
	"9":"wxyz",
}
//典型回溯题
//1. 回溯终止条件  拼接的字符串等于最终的字符串
//2. 每步做的事情  从对应字典中选出对应的字符串中的一位数值进行组合
var res = make([]string,0)
func letterCombinations(digits string) []string {
	res = res[0:0]
	if digits == "" {
		return res
	}
	backtracing("",digits)
	return res
}

func backtracing(currentLetter string,digits string)  {
	if len(currentLetter) == len(digits) {
		//截至条件
		res = append(res,currentLetter)
		return
	}
	//还未达到现在的长度
	//此时的已有字符串长度,，即为下一位字符对应的长度
	index := len(currentLetter)
	chars := dict[digits[index:index+1]]
	for i:=0;i<len(chars);i++ {
		//递归调用
		tempchar := chars[i:i+1]
		backtracing(currentLetter+tempchar,digits)
		//去除刚才添加的末位
		currentLetter = currentLetter[:index]
	}
}
