package mongodb

import (
	"context"

	"github.com/rekhin/generator/repository"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (c *Repository) Read(context.Context, []repository.Entity) error {
	return nil
}

func (c *Repository) Create(context.Context, []repository.Entity) error {
	return nil
}

func (c *Repository) Update(context.Context, []repository.Entity) error {
	return nil
}

func (c *Repository) Delete(context.Context, []repository.ID) error {
	return nil
}

func (c *Repository) SubscribeCreateUpdateFunc(context.Context, func(entities []repository.Entity)) error {
	return nil
}

func (c *Repository) SubscribeDeleteFunc(context.Context, func(ids []repository.ID)) error {
	return nil
}
