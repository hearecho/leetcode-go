package contest213

func canFormArray(arr []int, pieces [][]int) bool {
	index := make(map[int]int,0)
	for i:=0;i<len(pieces);i++ {
		index[pieces[i][0]] = i
	}
	for i:=0;i<len(arr); {
		temp := pieces[index[arr[i]]]
		for j:=0;j<len(temp);j++{
			if arr[i] != temp[j] {
				return false
			}
			i++
		}
	}
	return true
}
