package sedmax

import "github.com/rekhin/generator/repository"

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

func (o node) Equal(e repository.Entity) bool {
	second, ok := e.(*node)
	if !ok {
		return false
	}
	if o.id != second.id {
		return false
	}
	if o.parentID != second.parentID {
		return false
	}
	if o.name != second.name {
		return false
	}
	if o.sort != second.sort {
		return false
	}
	return true
}

func (o node) ID() repository.ID {
	return o.id
}

func (o node) ParentID() repository.ID {
	return o.parentID
}

func (o node) Name() string {
	return o.name
}

func (o node) Sort() int {
	return o.sort
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
