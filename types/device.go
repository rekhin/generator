package types

import (
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
)

type Device struct {
	Node
	Active bool
}

func (d Device) Equal(e repository.Entity) bool {
	if ok := d.Node.Equal(e); !ok {
		return false
	}
	second, ok := e.(Device)
	if !ok {
		return false
	}
	if d.GetActive() != second.GetActive() {
		return false
	}
	return true
}

func (d Device) GetActive() bool {
	return d.Active
}

func (d Device) Protocols() []sedmax.DeviceProtocol {
	return []sedmax.DeviceProtocol{}
}

var _ sedmax.Device = &Device{}
