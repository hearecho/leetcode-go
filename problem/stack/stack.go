package stack

//实现栈
type stack struct {
	data   []interface{}
	length int
}

func (s *stack) Push(e interface{}) bool {
	s.data = append(s.data, e)
	s.length++
	return true
}

func (s *stack) Pop() (interface{}, bool) {
	var e interface{}
	if s.length == 0 {
		return e,false
	}
	e = s.data[0]
	s.data = s.data[1:]
	return e,true
}
