package types

import (
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
)

type RTDTag struct {
	Tag
}

func (t RTDTag) Equal(e repository.Entity) bool {
	if ok := t.Tag.Equal(e); !ok {
		return false
	}
	if _, ok := e.(RTDTag); !ok {
		return false
	}
	return true
}

var _ sedmax.RTDTag = &RTDTag{}
