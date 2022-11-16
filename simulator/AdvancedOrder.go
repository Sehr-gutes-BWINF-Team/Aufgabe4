package simulator

type AdvancedOrder struct {
	Order       *Order
	WaitingTime int
	Priority    int
}

func (o *AdvancedOrder) GetPriority() int {
	return o.Priority
}

func (o *AdvancedOrder) IncreasePriority() {
	o.Priority++
}

func (o *AdvancedOrder) GetWaitingTime() int {
	return o.WaitingTime
}

func (o *AdvancedOrder) IncreaseWaitingTime() {
	o.WaitingTime++
}
