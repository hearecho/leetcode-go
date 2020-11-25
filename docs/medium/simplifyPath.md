# 简化路径
> 以 Unix 风格给出一个文件的绝对路径，你需要简化它。或者换句话说，将其转换为规范路径。
>在 Unix 风格的文件系统中，一个点（.）表示当前目录本身；此外，两个点 （..） 表示将目录切换到上一级（指向父目录

### 解题思路
> 使用栈进行解决,先将字符串转换为字符串数组
>

```go

func simplifyPath(path string) string {
	tempRes := ""
	strs := strings.Split(path,"/")
	stack := make([]string,0)
	for i:=0;i<len(strs);i++ {
		cur := strs[i]
		if cur == "" {
			continue
		}
		//当前显示上一级
		if cur == ".." {
			//stack中有元素,则取出
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur !="." && len(cur)> 0 {
			stack = append(stack,cur)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	tempRes = strings.Join(stack,"/")
	return "/"+tempRes
}
```