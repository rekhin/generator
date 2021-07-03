package inmemory

import (
	"context"
	"fmt"

	"github.com/rekhin/generator/repository"
)

type Repository struct {
	entities   map[repository.ID]repository.Entity // TODO make it safe
	createFunc func(entities []repository.Entity)
	updateFunc func(entities []repository.Entity)
	deleteFunc func(ids []repository.ID)
}

func NewRepository() *Repository {
	return &Repository{
		entities:   make(map[repository.ID]repository.Entity),
		createFunc: func(entities []repository.Entity) {},
		deleteFunc: func(ids []repository.ID) {},
	}
}

func (r *Repository) Read(_ context.Context) ([]repository.Entity, error) {
	var entities []repository.Entity
	for _, entity := range r.entities {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (r *Repository) Create(_ context.Context, entities []repository.Entity) error {
	var createEntities []repository.Entity
	for _, entity := range entities {
		id := entity.ID()
		if _, ok := r.entities[id]; ok {
			return fmt.Errorf("entity with id '%v' already exist", id)
		}
		createEntities = append(createEntities, entity)
	}
	for _, entity := range createEntities {
		id := entity.ID()
		r.entities[id] = entity
	}
	r.createFunc(createEntities)
	return nil
}

func (r *Repository) Update(_ context.Context, entities []repository.Entity) error {
	var updateEntities []repository.Entity
	for _, entity := range entities {
		id := entity.ID()
		if _, ok := r.entities[id]; !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		if r.entities[id].Equal(entity) {
			continue
		}
		updateEntities = append(updateEntities, entity)
	}
	for _, entity := range updateEntities {
		id := entity.ID()
		r.entities[id] = entity
	}
	r.updateFunc(updateEntities)
	return nil
}

func (r *Repository) Delete(_ context.Context, ids []repository.ID) error {
	var deleteIDs []repository.ID
	for _, id := range ids {
		if _, ok := r.entities[id]; !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		deleteIDs = append(deleteIDs, id)
	}
	for _, id := range deleteIDs {
		delete(r.entities, id)
	}
	r.deleteFunc(deleteIDs)
	return nil
}

func (r *Repository) SubscribeCreateFunc(_ context.Context, f func(entities []repository.Entity)) error {
	r.createFunc = f
	return nil
}

func (r *Repository) SubscribeUpdateFunc(_ context.Context, f func(entities []repository.Entity)) error {
	r.updateFunc = f
	return nil
}

func (r *Repository) SubscribeDeleteFunc(_ context.Context, f func(ids []repository.ID)) error {
	r.deleteFunc = f
	return nil
}

var _ repository.Reader = &Repository{}
var _ repository.Creator = &Repository{}
var _ repository.Updater = &Repository{}
var _ repository.Deleter = &Repository{}
var _ repository.CreateSubscriber = &Repository{}
var _ repository.UpdateSubscriber = &Repository{}
var _ repository.DeleteSubscriber = &Repository{}
