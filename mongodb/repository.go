package mongodb

import (
	"context"

	"github.com/rekhin/generator/repository"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (c *Repository) ReadEntities(context.Context) ([]repository.Entity, error) {
	return nil, nil
}

func (c *Repository) CreateEntities(context.Context, ...repository.Entity) error {
	return nil
}

func (c *Repository) UpdateEntities(context.Context, ...repository.Entity) error {
	return nil
}

func (c *Repository) DeleteEntitiesWithIDs(context.Context, ...repository.ID) error {
	return nil
}

func (c *Repository) SubscribeDelta(context.Context, func(repository.Delta)) error {
	return nil
}
