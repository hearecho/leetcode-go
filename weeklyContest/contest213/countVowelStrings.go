package contest213

var chars = []string{"a","e","i","o","u"}
func countVowelStrings(n int) int {
	res := 0
	backtrack("",&res,n,0)
	return res
}

func backtrack(custr string,res *int,l int,index int)  {
	if len(custr) == l {
		//fmt.Println(custr)
		*res++
		return
	}
	for i:=0;i<len(chars);i++ {
		if index <= i {
			custr = custr+chars[i]
			backtrack(custr,res,l,i)
			custr = custr[:len(custr)-1]
		}
	}
}