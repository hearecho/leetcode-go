# 二进制求和
> 给你两个二进制字符串，返回它们的和（用二进制表示）。输入为 非空 字符串且只包含数字 1 和 0。

### 解题思路
> 可以先将两个字符串补齐，使用0补齐，之后在进行相加，不过最后需要判断进位的情况

```go
func addBinary(a string, b string) string {
	s := ""
	var ex uint8
	diff := math.Abs(float64(len(a) - len(b)))
	for diff > 0{
		if len(a) > len(b) {
			b = "0" + b
		} else {
			a = "0" + a
		}
		diff--
	}

	i := len(a) - 1
	for i>=0{
		anum := a[i]-'0'
		bnum := b[i] - '0'
		t := anum + bnum + ex
		ex = t / 2
		t = t%2
		s = strconv.Itoa(int(t)) +s
		i--
	}
	if ex != 0 {
		s = strconv.Itoa(int(ex)) +s
	}
	return s
}
```