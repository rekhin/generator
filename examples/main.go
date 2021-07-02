package main

import (
	"context"
	"log"

	"github.com/rekhin/generator/configurator"
	"github.com/rekhin/generator/inmemory"
	"github.com/rekhin/generator/sedmax"
)

func main() {
	c := inmemory.NewConfigurator()

	ctx := context.TODO()

	if err := c.SubscribeCreateFunc(ctx, func(entities []configurator.Entity) {
		for _, e := range entities {
			log.Printf("created entity %v", e)
		}
	}); err != nil {
		log.Fatalf("subscribe create update func failed: %s", err)
	}

	if err := c.SubscribeUpdateFunc(ctx, func(entities []configurator.Entity) {
		for _, e := range entities {
			log.Printf("updated entity %v", e)
		}
	}); err != nil {
		log.Fatalf("subscribe create update func failed: %s", err)
	}

	if err := c.SubscribeDeleteFunc(ctx, func(ids []configurator.ID) {
		for _, id := range ids {
			log.Printf("deleted entity with id %v", id)
		}
	}); err != nil {
		log.Fatalf("subscribe create update func failed: %s", err)
	}

	if err := c.Create(ctx, []configurator.Entity{
		sedmax.NewObject(1, sedmax.RootObjectID, "New first object", 0),
		sedmax.NewObject(777, sedmax.RootObjectID, "New object", 0),
	}); err != nil {
		log.Fatalf("create failed: %s", err)
	}

	if err := c.Update(ctx, []configurator.Entity{
		sedmax.NewObject(777, sedmax.RootObjectID, "New object", 0),
	}); err != nil {
		log.Fatalf("update failed: %s", err)
	}

	if err := c.Update(ctx, []configurator.Entity{
		sedmax.NewObject(777, sedmax.RootObjectID, "My object", 0),
	}); err != nil {
		log.Fatalf("update failed: %s", err)
	}

	if err := c.Delete(ctx, []configurator.ID{
		sedmax.ObjectID(777),
	}); err != nil {
		log.Fatalf("delete failed: %s", err)
	}

	entities, err := c.Read(ctx)
	if err != nil {
		log.Fatalf("read failed: %s", err)
	}
	for _, e := range entities {
		log.Printf("readed entity %v", e)
	}
}
