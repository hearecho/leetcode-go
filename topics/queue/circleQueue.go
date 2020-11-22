package queue

type MyCircularQueue struct {
	nums   []int //存储
	length int   //总长度
	head   int   //队列头
	tail   int   // 队列尾
}

/** 初始化循环队列，长度为0 */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		nums: make([]int,k),
		length: k,
		head: -1,
		tail: -1,
	}
}

/** 插入元素 */
func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	if q.IsEmpty() {
		q.head = 0
	}
	//队尾插入
	q.tail = (q.tail+1)%q.length
	q.nums[q.tail] = value
	return true
}

/** 队头出队 */
func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	if q.head == q.tail {
		q.tail,q.head = -1,-1
		return true
	}
	//从对头取出
	q.head = (q.head+1)%q.length
	return true
}

/** 获取队列对头元素 */
func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.nums[q.head]
}

/** Get the last item from the queue. */
func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.nums[q.tail]
}

/** 判断队列是否为空 */
func (q *MyCircularQueue) IsEmpty() bool {
	return q.head == -1
}

/** 判断队列是否满了 */
func (q *MyCircularQueue) IsFull() bool {
	return q.head == ((q.tail+1)%q.length)
}
