package dailyProblem

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)
	gas = append(gas,gas...)
	oil := 0
	for i :=0;i<l;i++ {
		flag := true
	    oil = gas[i]
	    for j:=i+1;j<i+l+1;j++ {
		  temp := (j-1) % l
		  if oil < cost[temp] {
		  	flag = false
			 break
		  }
		  oil = oil - cost[temp] + gas[j]
	    }
	    if flag {
	    	return i
		}
	}
	return -1
}
