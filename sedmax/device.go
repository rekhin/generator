package sedmax

import "github.com/rekhin/generator/repository"

type Device interface {
	Node
	Active() bool
	Protocols() []DeviceProtocol
}

type DeviceProtocol interface {
	ProtocolID() int
	Name() string
	Active() bool
}

type DeviceID int

func (id DeviceID) Category() Category {
	return CategoryDevice
}

type device struct {
	node
	active bool
}

func NewDevice(id, parentID DeviceID, name string, sort int) Device {
	return &device{
		node: node{
			id:       id,
			parentID: parentID,
			name:     name,
			sort:     sort,
		},
	}
}

func (d device) Equal(e repository.Entity) bool {
	second, ok := e.(*device)
	if !ok {
		return false
	}
	if d.active != second.active {
		return false
	}
	return true
}

func (d device) Active() bool {
	return d.active
}

func (d device) Protocols() []DeviceProtocol {
	return []DeviceProtocol{}
}
