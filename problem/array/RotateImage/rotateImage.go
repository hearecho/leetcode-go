package RotateImage

func rotate(matrix [][]int) {
	//上下互换
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			swap(&matrix[i][j], &matrix[len(matrix)-i-1][j])
		}
	}
	//对角线对称互换
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			swap(&matrix[i][j], &matrix[j][i])
		}
	}
}

func swap(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
