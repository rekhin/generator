package repository

type Delta struct {
	CreateEntities        []Entity
	UpdateEntities        []Entity
	DeleteEntitiesWithIDs []ID
}
