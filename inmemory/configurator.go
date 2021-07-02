package inmemory

import (
	"context"
	"fmt"

	"github.com/rekhin/generator/configurator"
)

type Configurator struct {
	entities   map[configurator.ID]configurator.Entity // TODO make it safe
	createFunc func(entities []configurator.Entity)
	updateFunc func(entities []configurator.Entity)
	deleteFunc func(ids []configurator.ID)
}

func NewConfigurator() *Configurator {
	return &Configurator{
		entities:   make(map[configurator.ID]configurator.Entity),
		createFunc: func(entities []configurator.Entity) {},
		deleteFunc: func(ids []configurator.ID) {},
	}
}

func (c *Configurator) Read(_ context.Context) ([]configurator.Entity, error) {
	var entities []configurator.Entity
	for _, entity := range c.entities {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (c *Configurator) Create(_ context.Context, entities []configurator.Entity) error {
	var createEntities []configurator.Entity
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

func (c *Configurator) Update(_ context.Context, entities []configurator.Entity) error {
	var updateEntities []configurator.Entity
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

func (c *Configurator) Delete(_ context.Context, ids []configurator.ID) error {
	var deleteIDs []configurator.ID
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

func (c *Configurator) SubscribeCreateFunc(_ context.Context, f func(entities []configurator.Entity)) error {
	c.createFunc = f
	return nil
}

func (c *Configurator) SubscribeUpdateFunc(_ context.Context, f func(entities []configurator.Entity)) error {
	c.updateFunc = f
	return nil
}

func (c *Configurator) SubscribeDeleteFunc(_ context.Context, f func(ids []configurator.ID)) error {
	c.deleteFunc = f
	return nil
}

var _ configurator.Reader = &Configurator{}
var _ configurator.Creator = &Configurator{}
var _ configurator.Updater = &Configurator{}
var _ configurator.Deleter = &Configurator{}
var _ configurator.CreateSubscriber = &Configurator{}
var _ configurator.UpdateSubscriber = &Configurator{}
var _ configurator.DeleteSubscriber = &Configurator{}
