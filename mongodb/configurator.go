package mongodb

import (
	"context"

	"github.com/rekhin/generator/configurator"
)

type Configurator struct{}

func NewConfigurator() *Configurator {
	return &Configurator{}
}

func (c *Configurator) Read(context.Context, []configurator.Entity) error {
	return nil
}

func (c *Configurator) Create(context.Context, []configurator.Entity) error {
	return nil
}

func (c *Configurator) Update(context.Context, []configurator.Entity) error {
	return nil
}

func (c *Configurator) Delete(context.Context, []configurator.ID) error {
	return nil
}

func (c *Configurator) SubscribeCreateUpdateFunc(context.Context, func(entities []configurator.Entity)) error {
	return nil
}

func (c *Configurator) SubscribeDeleteFunc(context.Context, func(ids []configurator.ID)) error {
	return nil
}
