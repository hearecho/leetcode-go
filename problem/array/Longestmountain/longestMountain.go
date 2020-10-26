package Longestmountain

func longestMountain(a []int) (ans int) {
	n := len(a)
	left := make([]int, n)
	for i := 1; i < n; i++ {
		if a[i-1] < a[i] {
			left[i] = left[i-1] + 1
		}
	}
	right := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		if a[i+1] < a[i] {
			right[i] = right[i+1] + 1
		}
	}
	for i, l := range left {
		r := right[i]
		if l > 0 && r > 0 && l+r+1 > ans {
			ans = l + r + 1
		}
	}
	return
}
