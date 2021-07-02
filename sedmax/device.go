package sedmax

type Device interface {
	Node
	Active() bool
	Protocols() []Protocol
}

type Protocol interface {
	ID() int
	Name() string
	Active() bool
}

type DeviceID int

func (id DeviceID) Category() Category {
	return CategoryDevice
}
