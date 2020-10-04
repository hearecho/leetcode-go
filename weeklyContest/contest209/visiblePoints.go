package contest209

import (
	"math"
	"sort"
)

//关键:
//1. 在同一条直线上的点不影响，算作两个点
//2. 视野范围，假设人逆时针自传了d度，那么视野范围就是[d - angle/2, d + angle/2]，d的变化范围就是0~360
//3. 可以将点转换为角度，与location向东方向的角度,要转换为360度形式,
//4. 可能会出现在x轴上下的部分，所以需要在加上360，构成两圈
func visiblePoints(points [][]int, angle int, location []int) int {
	same := 0
	//转换
	angles := []float64{}
	a := 0.0
	for _,point := range points {
		x := point[0] - location[0]
		y := point[1] - location[1]
		if x == 0 && y==0 {
			same++
			continue
		}
		a = math.Atan2(float64(y), float64(x)) *180 /math.Pi
		if a<0 {
			a = a+360
		}
		angles = append(angles,a)
	}
	//排序
	sort.Float64s(angles)
	length := len(angles)
	//在加一圈
	for _,angle := range angles {
		angles = append(angles,angle+360)
	}
	max := 0
	//双指针
	for l,r:=0,0;l<length;l++ {
		for r+1<len(angles) && (angles[r+1]-angles[l]) <= float64(angle)+1e-8 {
			r++
		}
		max = int(math.Max(float64(max), float64(r-l+1)))
	}
	return max+same
}
