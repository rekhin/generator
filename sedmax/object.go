package sedmax

type Object interface {
	Node
}

type ObjectID int

const RootObjectID = 0

func (id ObjectID) GetCategory() Category {
	return CategoryObject
}
