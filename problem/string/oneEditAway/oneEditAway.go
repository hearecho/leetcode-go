package oneEditAway

func oneEditAway(first string, second string) bool {
	// 比较长度
	// 长度相差大于1 则必定不可能
	// 长度相差等于1 则只可能通过删除一个字符或者是插入一个字符
	// 长度相差等于0 则直接挨个遍历，如果有多于两个位置不同则为false
	l1, l2 := len(first), len(second)
	if l1 < l2 {
		return oneEditAway(second, first)
	}
	diff := l1 - l2
	if diff > 1 {
		return false
	} else if diff == 1 {
		// 长度差距为1 只能通过删除和插入来进行修改 删除和插入是一个意思 只是处理的字符串不同
		i, offest := 0, 0
		for i < l1 {
			if first[i] != second[i+offest] {
				offest += 1
				if offest > 1 {
					return false
				}
			} else {
				i++
			}
		}
		return offest > 1
	} else if diff == 0 {
		diffCount := 0
		for i := 0; i < l1; i++ {
			if first[i] != second[i] {
				diffCount++
			}
		}
		return diffCount <= 1
	}

	return true
}
