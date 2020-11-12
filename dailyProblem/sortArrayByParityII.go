package dailyProblem
/*
给定一个非负整数数组A， A 中一半整数是奇数，一半整数是偶数。
对数组进行排序，以便当A[i] 为奇数时，i也是奇数；当A[i]为偶数时， i 也是偶数。
你可以返回任何满足上述条件的数组作为答案。
 */
/*
数组长度肯定是偶数个，然后肯定是成对错位，所以双指针找到不对位的直接交换
 */
func sortArrayByParityII(A []int) []int {
	for i, j := 0, 1; i < len(A); i += 2 {
		// A[i]为基数，而i一直为偶数
		if A[i]%2 == 1 {
			for A[j]%2 == 1 {
				j += 2
			}
			//交换
			A[i], A[j] = A[j], A[i]
		}
	}
	return A
}
