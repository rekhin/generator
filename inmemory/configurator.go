package inmemory

import (
	"context"
	"fmt"

	"github.com/rekhin/generator/configurator"
)

type Configurator struct {
	entities         map[configurator.ID]configurator.Entity // TODO make it safe
	createUpdateFunc func(entities []configurator.Entity)
	deleteFunc       func(ids []configurator.ID)
}

func NewConfigurator() *Configurator {
	return &Configurator{
		entities:         make(map[configurator.ID]configurator.Entity),
		createUpdateFunc: func(entities []configurator.Entity) {},
		deleteFunc:       func(ids []configurator.ID) {},
	}
}

func (c *Configurator) Read(_ context.Context, entities *[]configurator.Entity) error {
	for _, entity := range c.entities {
		*entities = append(*entities, entity)
	}
	return nil
}

func (c *Configurator) Create(_ context.Context, entities []configurator.Entity) error {
	var createEntities []configurator.Entity
	for _, entity := range entities {
		entityID := entity.ID()
		if _, ok := c.entities[entityID]; ok {
			return fmt.Errorf("entity with id '%v' already exist", entityID)
		}
		c.entities[entity.ID()] = entity
		createEntities = append(createEntities, entity)
	}
	c.createUpdateFunc(createEntities)
	return nil
}

func (c *Configurator) Update(_ context.Context, entities []configurator.Entity) error {
	var updateEntities []configurator.Entity
	for _, entity := range entities {
		id := entity.ID()
		if _, ok := c.entities[id]; !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		c.entities[id] = entity
		updateEntities = append(updateEntities, entity)
	}
	c.createUpdateFunc(updateEntities)
	return nil
}

func (c *Configurator) Delete(_ context.Context, ids []configurator.ID) error {
	var deleteIDs []configurator.ID
	for _, id := range ids {
		if _, ok := c.entities[id]; !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		delete(c.entities, id)
		deleteIDs = append(deleteIDs, id)
	}
	c.deleteFunc(deleteIDs)
	return nil
}

func (c *Configurator) SubscribeCreateUpdate(_ context.Context, f func(entities []configurator.Entity)) error {
	c.createUpdateFunc = f
	return nil
}

func (c *Configurator) SubscribeDelete(_ context.Context, f func(ids []configurator.ID)) error {
	c.deleteFunc = f
	return nil
}
