package types

import (
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
)

type Object struct {
	Node
}

func (o Object) Equal(e repository.Entity) bool {
	if ok := o.Node.Equal(e); !ok {
		return false
	}
	if _, ok := e.(Object); !ok {
		return false
	}
	return true
}

var _ sedmax.Object = &Object{}
