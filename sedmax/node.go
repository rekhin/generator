package sedmax

import (
	"github.com/rekhin/generator/repository"
)

type Node interface {
	repository.Entity
	GetParentID() repository.ID
	GetName() string
	GetSort() int
}

type NodeID interface {
	GetCategory() Category
}

type Category int

const (
	CategoryObject = iota
	CategoryDevice
	CategoryRTDTag
	CategoryElectroTag
	CategoryEnergyTag
)
