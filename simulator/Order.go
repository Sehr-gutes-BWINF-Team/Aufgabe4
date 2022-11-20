package simulator

type Order struct {
	ID                      int
	EntryTime               int
	RequiredTime            int
	TimeLeftUntilCompletion int
}

func (o *Order) GetCompletion() int {
	return o.TimeLeftUntilCompletion
}

func (o *Order) IsCompleted() bool {
	return o.TimeLeftUntilCompletion == 0
}

func (o *Order) WorkOn() {
	o.TimeLeftUntilCompletion--
}
