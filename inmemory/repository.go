package inmemory

import (
	"context"
	"fmt"

	"github.com/rekhin/generator/repository"
)

type Repository struct {
	entities          map[repository.ID]repository.Entity
	publishDeltaFuncs []func(repository.Delta)
}

func NewRepository() *Repository {
	return &Repository{
		entities: make(map[repository.ID]repository.Entity),
	}
}

func (r *Repository) ReadEntity(_ context.Context, id repository.ID) (repository.Entity, bool, error) {
	entity, ok := r.entities[id]
	return entity, ok, nil
}

func (r *Repository) ReadEntities(_ context.Context) ([]repository.Entity, error) {
	var entities []repository.Entity
	for _, entity := range r.entities {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (r *Repository) CreateEntities(_ context.Context, entities ...repository.Entity) error {
	var createEntities []repository.Entity
	for _, entity := range entities {
		id := entity.GetID()
		if _, ok := r.entities[id]; ok {
			return fmt.Errorf("entity with id '%v' already exist", id)
		}
		r.entities[id] = entity
		createEntities = append(createEntities, entity)
	}
	for _, publishDeltaFunc := range r.publishDeltaFuncs {
		publishDeltaFunc(repository.Delta{CreateEntities: createEntities})
	}
	return nil
}

func (r *Repository) UpdateEntities(_ context.Context, entities ...repository.Entity) error {
	var updateEntities []repository.Entity
	for _, entity := range entities {
		id := entity.GetID()
		exist, ok := r.entities[id]
		if !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		if entity.Equal(exist) {
			continue
		}
		r.entities[id] = entity
		updateEntities = append(updateEntities, entity)
	}
	for _, publishDeltaFunc := range r.publishDeltaFuncs {
		publishDeltaFunc(repository.Delta{UpdateEntities: updateEntities})
	}
	return nil
}

func (r *Repository) DeleteEntitiesWithIDs(_ context.Context, ids ...repository.ID) error {
	var deleteEntitiesWithIDs []repository.ID
	for _, id := range ids {
		if _, ok := r.entities[id]; !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		delete(r.entities, id)
		deleteEntitiesWithIDs = append(deleteEntitiesWithIDs, id)
	}
	for _, publishDeltaFunc := range r.publishDeltaFuncs {
		publishDeltaFunc(repository.Delta{DeleteEntitiesWithIDs: deleteEntitiesWithIDs})
	}
	return nil
}

func (r *Repository) PublishDelta(ctx context.Context, d repository.Delta) error {
	if err := r.CreateEntities(ctx, d.CreateEntities...); err != nil {
		return fmt.Errorf("create entities failed: %s", err)
	}
	if err := r.UpdateEntities(ctx, d.UpdateEntities...); err != nil {
		return fmt.Errorf("update entities failed: %s", err)
	}
	if err := r.DeleteEntitiesWithIDs(ctx, d.DeleteEntitiesWithIDs...); err != nil {
		return fmt.Errorf("delete entities with ids failed: %s", err)
	}
	return nil
}

func (r *Repository) SubscribeDelta(_ context.Context, f func(repository.Delta)) error {
	r.publishDeltaFuncs = append(r.publishDeltaFuncs, f)
	return nil
}

var (
	_ repository.EntityReader    = &Repository{}
	_ repository.EntitiesReader  = &Repository{}
	_ repository.EntitiesCreator = &Repository{}
	_ repository.EntitiesUpdater = &Repository{}
	_ repository.EntitiesDeleter = &Repository{}
	_ repository.DeltaPublisher  = &Repository{}
	_ repository.DeltaSubscriber = &Repository{}
)
