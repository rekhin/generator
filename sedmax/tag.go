package sedmax

type Tag interface {
	Node
	GetActive() bool
	GetProfile() Profile
	GetType() Type
	GetUnit() string
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
