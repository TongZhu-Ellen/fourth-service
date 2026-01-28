package main



type Outbox struct {


   
	Orderp *Order // uuid
	Action  Action
	Processed bool
	TryCount int

	
}

type Action string 
const (
	ReserveInvent Action = "RESERVE_INVENT"
	PayWithLuck   Action = "PAY_WITH_LUCK"
	ReleaseInvent Action = "RELEASE_INVENT"
)


















