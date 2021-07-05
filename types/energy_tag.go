package types

import (
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
)

type EnergyTag struct {
	Tag
}

func (t EnergyTag) Equal(e repository.Entity) bool {
	if ok := t.Tag.Equal(e); !ok {
		return false
	}
	if _, ok := e.(EnergyTag); !ok {
		return false
	}
	return true
}

var _ sedmax.EnergyTag = &EnergyTag{}
