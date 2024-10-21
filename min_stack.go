package livetest

type MinStack struct {
	data []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.data) == 0 {
		return 0
	}
	min := this.data[0]
	for _, v := range this.data {
		if v < min {
			min = v
		}
	}
	return min
}
