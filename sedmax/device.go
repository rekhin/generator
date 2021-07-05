package sedmax

type Device interface {
	Node
	GetActive() bool
	Protocols() []DeviceProtocol
}

type DeviceProtocol interface {
	GetProtocolID() int
	GetName() string
	GetActive() bool
}

type DeviceID int

func (id DeviceID) GetCategory() Category {
	return CategoryDevice
}
