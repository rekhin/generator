package configurator

type Entity interface {
	ID() ID
}

type ID string
