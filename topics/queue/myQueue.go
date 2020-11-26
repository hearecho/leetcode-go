package queue

/**
使用栈实现队列
*/
type MyQueue struct {
	//使用两个栈
	instack  []int
	outstack []int
}

/** Initialize your data structure here. */
func Constructor1() MyQueue {
	return MyQueue{
		instack:  make([]int, 0),
		outstack: make([]int, 0),
	}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.instack = append(q.instack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	//出队列，先从用于出队的栈中
	if len(q.outstack) != 0 {
		res := q.outstack[len(q.outstack)-1]
		q.outstack = q.outstack[:len(q.outstack)-1]
		return res
	}
	//用于出队的栈为空，所以将用于入队的栈中的元素全部出栈到用于出队的栈中
	if len(q.outstack) == 0 {
		for len(q.instack) != 0 {
			q.outstack = append(q.outstack, q.instack[len(q.instack)-1])
			q.instack = q.instack[:len(q.instack)-1]
		}
		return q.Pop()
	}
	return -1
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	//出队列，先从用于出队的栈中
	if len(q.outstack) != 0 {
		res := q.outstack[len(q.outstack)-1]
		return res
	}
	//用于出队的栈为空，所以将用于入队的栈中的元素全部出栈到用于出队的栈中
	if len(q.outstack) == 0 {
		for len(q.instack) != 0 {
			q.outstack = append(q.outstack, q.instack[len(q.instack)-1])
			q.instack = q.instack[:len(q.instack)-1]
		}
		return q.Peek()
	}
	return -1
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	if len(q.instack) == 0 && len(q.outstack) == 0 {
		return true
	}
	return false
}
