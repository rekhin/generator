package inmemory

import (
	"context"
	"fmt"

	"github.com/rekhin/generator/repository"
)

// TODO это и есть репозиторий
type Storage struct {
	entities map[repository.ID]repository.Entity
}

func NewStorage() *Storage {
	return &Storage{
		entities: make(map[repository.ID]repository.Entity),
	}
}

func (s *Storage) ReadEntity(_ context.Context, id repository.ID) (repository.Entity, bool, error) {
	entity, ok := s.entities[id]
	return entity, ok, nil
}

func (s *Storage) ReadEntities(_ context.Context) ([]repository.Entity, error) {
	var entities []repository.Entity
	for _, entity := range s.entities {
		entities = append(entities, entity)
	}
	return entities, nil
}

func (s *Storage) CreateEntities(_ context.Context, entities ...repository.Entity) error {
	for _, entity := range entities {
		id := entity.ID()
		if _, ok := s.entities[id]; ok {
			return fmt.Errorf("entity with id '%v' already exist", id)
		}
		s.entities[id] = entity
	}
	return nil
}

func (s *Storage) UpdateEntities(_ context.Context, entities ...repository.Entity) error {
	for _, entity := range entities {
		id := entity.ID()
		exist, ok := s.entities[id]
		if !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		if entity.Equal(exist) {
			continue
		}
		s.entities[id] = entity
	}
	return nil
}

func (s *Storage) DeleteEntitiesWithIDs(_ context.Context, ids ...repository.ID) error {
	for _, id := range ids {
		if _, ok := s.entities[id]; !ok {
			return fmt.Errorf("entity with id '%v' does not exist", id)
		}
		delete(s.entities, id)
	}
	return nil
}
