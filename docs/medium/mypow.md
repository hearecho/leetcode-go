# n次幂
> 实现 pow(x, n) ，即计算 x 的 n 次幂函数
>
### 注意点
1. x为浮点数且-100.0 < x < 100.0
2. n为32位有符号整数

### 解题思路
>使用递归，递归终止条件就是 n等于0或者n等于1的时候，分别返回1或者是x
> 每次对n除以2，便可以一直划分，
>如果最后n是奇数则应该多乘以x，因为1/2==0
```go
func myPow(x float64, n int) float64 {
	if n==0 {
		return 1
	}
	if n==1 {
		return x
	}
	if n < 0{
		n = -n
		x = 1/x
	}
	temp := myPow(x,n/2)
	if n % 2 == 0 {
		return temp *temp
	}
	return temp * temp *x
}
```