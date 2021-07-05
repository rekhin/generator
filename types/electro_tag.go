package types

import (
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
)

type ElecrtoTag struct {
	Tag
}

func (t ElecrtoTag) Equal(e repository.Entity) bool {
	if ok := t.Tag.Equal(e); !ok {
		return false
	}
	if _, ok := e.(ElecrtoTag); !ok {
		return false
	}
	return true
}

var _ sedmax.ElecrtoTag = &ElecrtoTag{}
