package sedmax

// ElecrtoTag ...
type ElecrtoTag interface {
	Tag
}

type ElecrtoTagID struct {
	DeviceID
	Profile
	Code string
}

func (id ElecrtoTagID) GetCategory() Category {
	return CategoryElectroTag
}
