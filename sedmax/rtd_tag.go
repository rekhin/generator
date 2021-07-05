package sedmax

// RTDTag ...
type RTDTag interface {
	Tag
}

type RTDTagID int

func (id RTDTagID) GetCategory() Category {
	return CategoryRTDTag
}
