# 旋转链表
> 给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。
>
### 注意点
> k有可能是大于链表长度的，所以需要在开始对k进行处理

### 解题思路
![ByH1rF.png](https://s1.ax1x.com/2020/11/03/ByH1rF.png)

```go
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	newHead := &ListNode{0,head}
	l := 0
	cur := newHead
	for cur.Next != nil {
		l++
		cur = cur.Next
	}
	k = k%l
	//如果k刚好是l的整数倍，那样的话，肯定会旋转为原来的样子,所以直接返回原字符串
	if k  == 0 {
		return head
	}
	cur.Next = head
	cur = newHead
	//找到中间
	for i := l - k; i > 0; i-- {
		cur = cur.Next
	}
	tempRes := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return tempRes.Next
}
```