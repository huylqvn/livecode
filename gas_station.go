package livetest

func canCompleteCircuit(gas []int, cost []int) int {
	start, total, tank := 0, 0, 0
	for i := 0; i < len(gas); i++ {
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			total += tank
			tank = 0
		}
	}
	if total+tank < 0 {
		return -1
	}
	return start
}
