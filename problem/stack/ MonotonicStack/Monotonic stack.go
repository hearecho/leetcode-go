package _MonotonicStack

// 单调栈

type Monotonstack []int

func (m *Monotonstack) Push(x int) {
	for len(*m) > 0 && (*m)[len(*m)-1] < x {
		*m = (*m)[:len(*m)-1]
	}
	*m = append(*m, x)
}

// 之所以传入值 是因为可能需要取出的值已经不存在了

func (m *Monotonstack) Pop(x int) {
	if x == (*m)[len(*m)-1] {
		*m = (*m)[:len(*m)-1]
	}
}
