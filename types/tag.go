package types

import (
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
)

type Tag struct {
	Node
	Active  bool
	Profile sedmax.Profile
	Type    sedmax.Type
	Unit    string
}

func (t Tag) Equal(e repository.Entity) bool {
	if ok := t.Node.Equal(e); !ok {
		return false
	}
	second, ok := e.(Tag)
	if !ok {
		return false
	}
	if t.GetActive() != second.GetActive() {
		return false
	}
	if t.GetProfile() != second.GetProfile() {
		return false
	}
	if t.GetType() != second.GetType() {
		return false
	}
	if t.GetUnit() != second.GetUnit() {
		return false
	}
	return true
}

var _ sedmax.Tag = &Tag{}

func (t Tag) GetActive() bool {
	return t.Active
}

func (t Tag) GetProfile() sedmax.Profile {
	return t.Profile
}

func (t Tag) GetType() sedmax.Type {
	return t.Type
}

func (t Tag) GetUnit() string {
	return t.Unit
}
