package inmemory

import (
	"context"
	"fmt"

	"github.com/rekhin/generator/repository"
)

type Repository struct {
	entities              map[repository.ID]repository.Entity // TODO make it safe
	createEntities        []repository.Entity                 // TODO make it safe
	updateEntities        []repository.Entity                 // TODO make it safe
	deleteEntitiesWithIDs []repository.ID                     // TODO make it safe
	deltaFuncs            []func(repository.Delta)
}

func NewRepository() *Repository {
	return &Repository{
		entities: make(map[repository.ID]repository.Entity),
	}
}

func (r *Repository) ReadEntities(_ context.Context) ([]repository.Entity, error) {
	var entities []repository.Entity
	for _, entity := range r.entities {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (r *Repository) CreateEntities(_ context.Context, entities ...repository.Entity) error {
	for _, entity := range entities {
		id := entity.GetID()
		if _, ok := r.entities[id]; ok {
			return fmt.Errorf("entity with id '%v' already exist", id)
		}
		r.createEntities = append(r.createEntities, entity)
	}
	return nil
}

func (r *Repository) UpdateEntities(_ context.Context, entities ...repository.Entity) error {
	for _, entity := range entities {
		id := entity.GetID()
		exist, ok := r.entities[id]
		if !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		if entity.Equal(exist) {
			continue
		}
		r.updateEntities = append(r.updateEntities, entity)
	}
	return nil
}

func (r *Repository) DeleteEntitiesWithIDs(_ context.Context, ids ...repository.ID) error {
	for _, id := range ids {
		if _, ok := r.entities[id]; !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		r.deleteEntitiesWithIDs = append(r.deleteEntitiesWithIDs, id)
	}
	return nil
}

func (r *Repository) Accept(ctx context.Context) error {
	for _, entity := range r.createEntities {
		id := entity.GetID()
		r.entities[id] = entity
	}
	for _, entity := range r.updateEntities {
		id := entity.GetID()
		r.entities[id] = entity
	}
	for _, id := range r.deleteEntitiesWithIDs {
		delete(r.entities, id)
	}
	for _, deltaFunc := range r.deltaFuncs {
		deltaFunc(repository.Delta{
			CreateEntities:        r.createEntities,
			UpdateEntities:        r.updateEntities,
			DeleteEntitiesWithIDs: r.deleteEntitiesWithIDs,
		})
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
	if err := r.Accept(ctx); err != nil {
		return fmt.Errorf("accept failed: %s", err)
	}
	return nil
}

func (r *Repository) SubscribeDelta(_ context.Context, f func(repository.Delta)) error {
	r.deltaFuncs = append(r.deltaFuncs, f)
	return nil
}

var (
	_ repository.EntitiesReader  = &Repository{}
	_ repository.EntitiesCreator = &Repository{}
	_ repository.EntitiesUpdater = &Repository{}
	_ repository.EntitiesDeleter = &Repository{}
	_ repository.DeltaPublisher  = &Repository{}
	_ repository.DeltaSubscriber = &Repository{}
)
