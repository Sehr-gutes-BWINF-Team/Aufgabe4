package simulator

type Order struct {
	Entry      int
	Time       int
	Completion int
}

func (o *Order) getCompletion() int {
	return o.Completion
}

func (o *Order) isCompleted() bool {
	return o.Completion == 0
}

func (o *Order) workOn() {
	o.Completion--
}
