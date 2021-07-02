package main

import (
	"context"
	"log"

	"github.com/rekhin/generator/configurator"
	"github.com/rekhin/generator/inmemory"
	"github.com/rekhin/generator/sedmax"
)

func main() {
	ctx := context.TODO()
	c := inmemory.NewConfigurator()

	if err := c.SubscribeCreateUpdateFunc(ctx, func(entities []configurator.Entity) {
		for _, e := range entities {
			log.Printf("created or updated entity %v", e)
		}
	}); err != nil {
		log.Fatalf("subscribe create update func failed: %s", err)
	}

	if err := c.Create(ctx, []configurator.Entity{sedmax.NewObject(777, sedmax.RootObjectID, "New object", 0)}); err != nil {
		log.Fatalf("create failed: %s", err)
	}

	if err := c.Update(ctx, []configurator.Entity{sedmax.NewObject(777, sedmax.RootObjectID, "My object", 0)}); err != nil {
		log.Fatalf("update failed: %s", err)
	}
}
