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

func (c *Repository) Read(_ context.Context) ([]repository.Entity, error) {
	var entities []repository.Entity
	for _, entity := range c.entities {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (c *Repository) Create(_ context.Context, entities []repository.Entity) error {
	var createEntities []repository.Entity
	for _, entity := range entities {
		id := entity.ID()
		if _, ok := c.entities[id]; ok {
			return fmt.Errorf("entity with id '%v' already exist", id)
		}
		createEntities = append(createEntities, entity)
	}
	for _, entity := range createEntities {
		id := entity.ID()
		c.entities[id] = entity
	}
	c.createFunc(createEntities)
	return nil
}

func (c *Repository) Update(_ context.Context, entities []repository.Entity) error {
	var updateEntities []repository.Entity
	for _, entity := range entities {
		id := entity.ID()
		if _, ok := c.entities[id]; !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		if c.entities[id].Equal(entity) {
			continue
		}
		updateEntities = append(updateEntities, entity)
	}
	for _, entity := range updateEntities {
		id := entity.ID()
		c.entities[id] = entity
	}
	c.updateFunc(updateEntities)
	return nil
}

func (c *Repository) Delete(_ context.Context, ids []repository.ID) error {
	var deleteIDs []repository.ID
	for _, id := range ids {
		if _, ok := c.entities[id]; !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		deleteIDs = append(deleteIDs, id)
	}
	for _, id := range deleteIDs {
		delete(c.entities, id)
	}
	c.deleteFunc(deleteIDs)
	return nil
}

func (c *Repository) SubscribeCreateFunc(_ context.Context, f func(entities []repository.Entity)) error {
	c.createFunc = f
	return nil
}

func (c *Repository) SubscribeUpdateFunc(_ context.Context, f func(entities []repository.Entity)) error {
	c.updateFunc = f
	return nil
}

func (c *Repository) SubscribeDeleteFunc(_ context.Context, f func(ids []repository.ID)) error {
	c.deleteFunc = f
	return nil
}

var _ repository.Reader = &Repository{}
var _ repository.Creator = &Repository{}
var _ repository.Updater = &Repository{}
var _ repository.Deleter = &Repository{}
var _ repository.CreateSubscriber = &Repository{}
var _ repository.UpdateSubscriber = &Repository{}
var _ repository.DeleteSubscriber = &Repository{}
