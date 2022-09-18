package simulator

type Order struct {
	Id         int
	Entry      int
	Time       int
	Completion int
}

func (o *Order) GetCompletion() int {
	return o.Completion
}

func (o *Order) IsCompleted() bool {
	return o.Completion == 0
}

func (o *Order) WorkOn() {
	o.Completion--
}
