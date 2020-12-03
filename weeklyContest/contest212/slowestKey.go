package contest210
func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxIndex := 0
	maxLast := releaseTimes[0]
	for i:=1;i<len(keysPressed);i++ {
		last := releaseTimes[i] - releaseTimes[i-1]
		if last > maxLast {
			maxLast = last
			maxIndex = i
		}
		if last == maxLast && keysPressed[maxIndex] < keysPressed[i]{
			maxIndex = i
		}
	}
	return keysPressed[maxIndex]
}