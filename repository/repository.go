package repository

import "context"

type EntityReader interface {
	ReadEntity(context.Context, ID) (Entity, bool, error)
}

type EntitiesReader interface {
	ReadEntities(context.Context) ([]Entity, error)
}

type EntitiesCreator interface {
	CreateEntities(context.Context, ...Entity) error
}

type EntitiesUpdater interface {
	UpdateEntities(context.Context, ...Entity) error
}

type EntitiesDeleter interface {
	DeleteEntitiesWithIDs(context.Context, ...ID) error
}

type DeltaPublisher interface {
	PublishDelta(context.Context, Delta) error
}

type DeltaSubscriber interface {
	SubscribeDelta(context.Context, func(Delta)) error
}
