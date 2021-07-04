package inmemory

import (
	"context"
	"fmt"

	"github.com/rekhin/generator/repository"
)

// TODO move to package cache
type Cache struct {
	entityReader          repository.EntityReader
	createEntities        []repository.Entity // TODO make it safe
	updateEntities        []repository.Entity // TODO make it safe
	deleteEntitiesWithIDs []repository.ID     // TODO make it safe
}

func NewCache(r repository.EntityReader) *Cache {
	return &Cache{
		entityReader: r,
	}
}

func (c *Cache) CreateEntities(ctx context.Context, entities ...repository.Entity) error {
	for _, entity := range entities {
		id := entity.ID()
		_, ok, err := c.entityReader.ReadEntity(ctx, id)
		if err != nil {
			return fmt.Errorf("read entity failed: %s", err)
		}
		if ok {
			return fmt.Errorf("entity with id '%v' already exist", id)
		}
		c.createEntities = append(c.createEntities, entity)
	}
	return nil
}

func (c *Cache) UpdateEntities(ctx context.Context, entities ...repository.Entity) error {
	for _, entity := range entities {
		id := entity.ID()
		exist, ok, err := c.entityReader.ReadEntity(ctx, id)
		if err != nil {
			return fmt.Errorf("read entity failed: %s", err)
		}
		if !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		if entity.Equal(exist) {
			continue
		}
		c.updateEntities = append(c.updateEntities, entity)
	}
	return nil
}

func (c *Cache) DeleteEntitiesWithIDs(ctx context.Context, ids ...repository.ID) error {
	for _, id := range ids {
		_, ok, err := c.entityReader.ReadEntity(ctx, id)
		if err != nil {
			return fmt.Errorf("read entity failed: %s", err)
		}
		if !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		c.deleteEntitiesWithIDs = append(c.deleteEntitiesWithIDs, id)
	}
	return nil
}

func (c *Cache) GetDelta(ctx context.Context) repository.Delta {
	// if err := c.storage.CreateEntities(ctx, c.createEntities...); err != nil {
	// 	return repository.Delta{}, fmt.Errorf("create entities failed: %s", err)
	// }
	// if err := c.storage.UpdateEntities(ctx, c.updateEntities...); err != nil {
	// 	return repository.Delta{}, fmt.Errorf("update entities failed: %s", err)
	// }
	// if err := c.storage.DeleteEntitiesWithIDs(ctx, c.deleteEntitiesWithIDs...); err != nil {
	// 	return repository.Delta{}, fmt.Errorf("delete entities with ids failed: %s", err)
	// }

	return repository.Delta{
		CreateEntities:        c.createEntities,
		UpdateEntities:        c.updateEntities,
		DeleteEntitiesWithIDs: c.deleteEntitiesWithIDs,
	}
}

func (c *Cache) Clear(ctx context.Context) {
	c.createEntities = []repository.Entity{}
	c.updateEntities = []repository.Entity{}
	c.deleteEntitiesWithIDs = []repository.ID{}
}
