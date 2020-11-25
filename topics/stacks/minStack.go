package stacks

type MinStack struct {
	nums    []int
	minNums []int
}

func Constructor() MinStack {
	return MinStack{
		nums:    make([]int, 0),
		minNums: make([]int, 0),
	}
}

func (s *MinStack) Push(x int) {
	//插入的时候，直接插入到nums中，然后比较minNums中的数据
	s.nums = append(s.nums, x)
	//将此时最小的值插入到minNums中
	if len(s.minNums) == 0 {
		s.minNums = append(s.minNums, x)
	} else {
		temp := s.minNums[len(s.minNums)-1]
		if temp > x {
			s.minNums = append(s.minNums, x)
		} else {
			s.minNums = append(s.minNums, temp)
		}
	}
}

func (s *MinStack) Pop() {
	s.nums = s.nums[:len(s.nums)-1]
	s.minNums = s.minNums[:len(s.minNums)-1]
}

func (s *MinStack) Top() int {
	return s.nums[len(s.nums)-1]
}

func (s *MinStack) GetMin() int {
	return s.minNums[len(s.minNums)-1]
}
