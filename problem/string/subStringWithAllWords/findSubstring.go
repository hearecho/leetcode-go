package subStringWithAllWords

/**
字符串数组中字符串长度相同,数组中的字符串拼接顺序不同
使用到哈希表
 */
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	res := []int{}
	counter := map[string]int{}
	//计数
	for _,w := range words {
		counter[w]++
	}
	//totallen为数组中字符串拼接之后的最长长度  
	length,totalLen,tmpCounter := len(words[0]),len(words[0])*len(words),copyMap(counter)
	//由于length固定，所以可以i进行固定步长
	//start为最后结果的起始,
	for i,start :=0,0;i<len(s)-length+1 && start <len(s)-length+1;i++ {
		if tmpCounter[s[i:i+length]] > 0{
			//减少相应单词计数
			tmpCounter[s[i:i+length]]--
			if checkWords(tmpCounter) && (i+length-start)==totalLen {
				res = append(res,start)
				continue
			}
			i = i+length-1
		} else {
			start++
			i = start-1
			//复制word计数
			tmpCounter = copyMap(counter)
		}

	}
	return res
}
/**
检查counter中是否还有剩余单词
 */
func checkWords(counter map[string]int) bool{
	flag := true
	for _,v :=  range counter {
		if v >0 {
			flag = false
			break
		}
	}
	return flag
}

func copyMap(counter map[string]int) map[string]int {
	c := map[string]int{}
	for k,v := range counter {
		c[k] = v
	}
	return c
}
