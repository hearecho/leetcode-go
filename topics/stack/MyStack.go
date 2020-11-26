package stack

type MyStack struct {
	queue1 []int
	queue2 []int
}

/** Initialize your data structure here. */
func Constructor1() MyStack {
	return MyStack{
		queue1: make([]int, 0),
		queue2: make([]int, 0),
	}
}

/** Push element x onto stack. */
func (s *MyStack) Push(x int) {
	//哪个不为空放在哪个里面
	if len(s.queue1) != 0 {
		s.queue1 = append(s.queue1, x)
	} else {
		s.queue2 = append(s.queue2, x)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (s *MyStack) Pop() int {
	//队列是FIFO，而且是使用队列实现栈，那么就要使用队列的规则
	//将第一个队列除了最后一个元素外元素放入第二个队列中里面
	if len(s.queue1) != 0 {
		for i := 0; i < len(s.queue1)-1; i++ {
			//出队列1入队列2
			s.queue2 = append(s.queue2, s.queue1[i])
		}
		//截断
		temp := s.queue1[len(s.queue1)-1]
		s.queue1 = s.queue1[len(s.queue1):]
		return temp
	}
	if len(s.queue2) != 0 {
		for i := 0; i < len(s.queue2)-1; i++ {
			//出队列1入队列2
			s.queue1 = append(s.queue1, s.queue2[i])
		}
		//截断
		temp := s.queue2[len(s.queue2)-1]
		s.queue2 = s.queue2[len(s.queue2):]
		return temp
	}
	return -1
}

/** Get the top element. */
func (s *MyStack) Top() int {
	if len(s.queue1) != 0 {
		return s.queue1[len(s.queue1)-1]
	}
	if len(s.queue2) != 0 {
		return s.queue2[len(s.queue2)-1]
	}
	return -1
}

/** Returns whether the stack is empty. */
func (s *MyStack) Empty() bool {
	if len(s.queue2) == 0 && len(s.queue1) == 0 {
		return true
	}
	return false
}
