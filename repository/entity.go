package repository

type Entity interface {
	Equal(Entity) bool
	GetID() ID
}

type ID interface{}
