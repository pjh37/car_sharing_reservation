package datamodel

type Reservation struct {
	ID       string
	UserID   string
	UserName string
	Address  string
	Position Position
	reserved bool
}
type Position struct {
	latitude  float64
	longitude float64
}

func (Reservation) CollectionName() string {
	return "reservation"
}
