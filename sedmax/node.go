package sedmax

import "github.com/rekhin/generator/configurator"

type Node interface {
	configurator.Entity
	ParentID() configurator.ID
	Name() string
	Sort() int
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
