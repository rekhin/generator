package sedmax

import (
	"github.com/rekhin/generator/repository"
)

type Node interface {
	repository.Entity
	ParentID() repository.ID
	Name() string
	Sort() int
}

type node struct {
	id       repository.ID
	parentID repository.ID
	name     string
	sort     int
}

func (n node) Equal(e repository.Entity) bool {
	second, ok := e.(Node)
	if !ok {
		return false
	}
	if n.ID() != second.ID() {
		return false
	}
	if n.ParentID() != second.ParentID() {
		return false
	}
	if n.Name() != second.Name() {
		return false
	}
	if n.Sort() != second.Sort() {
		return false
	}
	return true
}

func (n node) ID() repository.ID {
	return n.id
}

func (n node) ParentID() repository.ID {
	return n.parentID
}

func (n node) Name() string {
	return n.name
}

func (n node) Sort() int {
	return n.sort
}

// type ID interface {
// 	Category() Category
// }

type Category int

const (
	CategoryObject = iota
	CategoryDevice
	CategoryRTDTag
	CategoryElectroTag
	CategoryEnergyTag
)
