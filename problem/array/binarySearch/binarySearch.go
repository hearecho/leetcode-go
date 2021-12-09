package binarySearch

func base(nums []int, target int) int {
	// 我们遵循找下界的方式 即每次的区间改为左闭右开
	// 所有的区间调整均要遵循左闭右开的原则
	l, h := 0, len(nums)
	for l < h {
		mid := l + (h-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			h = mid
		}
	}
	return -1
}

func base2(nums []int, target int) int {
	// 如果使用两边全部闭合的区间 则对应需要修改循环的结束条件以及每次区间修改
	// 包括区间的起始也要保证是两边闭合即均可访问到
	l, h := 0, len(nums)-1
	// 两边闭合的情况下会出现 l==h的情况
	for l <= h {
		mid := l + (h-l)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			//要保证两边均都要闭合
			h = mid - 1
		}
	}
	return -1
}

func searchLeft(nums []int, target int) int {
	// 寻找左侧边界问题
	// 应该使用左闭右开的区间大小
	l, h := 0, len(nums)
	for l < h {
		// 只有这样最后可以精确的返回左边界
		// 因为不是为了寻找确定的值的位置
		// 所以等于的情况并不会进行返回 也是修改区间
		// 可以看到 相等的情况和大于的情况是相同的处理，所以可以进行合并
		mid := l + (h-l)/2
		if nums[mid] == target {
			// 我们是寻找左边界，所以需要修改h
			// 又因为是左闭右开
			h = mid
		} else if nums[mid] < target {
			//TODO l = ...
			l = mid + 1
		} else if nums[mid] > target {
			//TODO h = .
			h = mid
		}
	}
	return l
}
