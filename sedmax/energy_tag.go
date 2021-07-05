package sedmax

// EnergyTag ...
type EnergyTag interface {
	Tag
}

type EnergyTagID int

func (id EnergyTagID) GetCategory() Category {
	return CategoryEnergyTag
}
