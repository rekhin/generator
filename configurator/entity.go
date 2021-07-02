package configurator

type Entity interface {
	Equal(Entity) bool
	ID() ID
}

type ID interface{}
