package sedmax

type Tag interface {
	Node
	Active() bool
	Profile() Profile
	Type() Type
	Unit() string
}

type Profile int

const (
	ProfileRTD = iota
	Profile1m
	Profile3m
	Profile15m
	Profile30m
	Profile1h
	Profile1d
	Profile1mon
)

type Type int

const (
	TypeBool    = iota // bool
	TypeInt8           // int8
	TypeInt16          // int16
	TypeInt32          // int32
	TypeInt64          // int64
	TypeUint8          // uint8
	TypeUint16         // uint16
	TypeUint32         // uint32
	TypeUint64         // uint64
	TypeFloat32        // float32
	TypeFloat64        // float64
	TypeString         // string
)

// RTDTag ...
type RTDTag interface {
	Tag
}

type RTDTagID int

func (id RTDTagID) Category() Category {
	return CategoryRTDTag
}

// ElecrtoTag ...
type ElecrtoTag interface {
	Tag
}

type ElecrtoTagID struct {
	DeviceID
	Profile
	Code string
}

func (id ElecrtoTagID) Category() Category {
	return CategoryElectroTag
}

// EnergyTag ...
type EnergyTag interface {
	Tag
}

type EnergyTagID int

func (id EnergyTagID) Category() Category {
	return CategoryEnergyTag
}
