package MaxQueue

/**
要求均摊时间复杂度为O(1)
单调队列
我们可以使用两个队列来操作，类似于最小栈
*/

type MaxQueue struct {
	Origin []int
	Max    []int
}

func Constructor() MaxQueue {
	return MaxQueue{Origin: []int{}, Max: []int{}}
}

func (q *MaxQueue) Max_value() int {
	if len(q.Origin) == 0 {
		return -1
	}
	return q.Max[0]
}

func (q *MaxQueue) Push_back(value int) {
	// 加入到队列的值,
	q.Origin = append(q.Origin, value)
	//对于Max 需要将队尾小于该数字的元素去除
	i := len(q.Max) - 1
	for ; i >= 0; i-- {
		if q.Max[i] > value {
			break
		}
	}
	q.Max = append(q.Max[:i+1], value)
}

func (q *MaxQueue) Pop_front() int {
	if len(q.Origin) == 0 {
		return -1
	}
	temp := q.Origin[0]
	if temp == q.Max[0] {
		q.Max = q.Max[1:]
	}
	q.Origin = q.Origin[1:]
	return temp
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
