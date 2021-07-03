package main

import (
	"context"
	"log"

	"github.com/rekhin/generator/inmemory"
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
)

func main() {
	objectRepository := inmemory.NewRepository()

	ctx := context.TODO()

	if err := objectRepository.SubscribeCreateFunc(ctx, func(entities []repository.Entity) {
		for _, e := range entities {
			log.Printf("created entity %+v", e)
		}
	}); err != nil {
		log.Fatalf("subscribe create update func failed: %s", err)
	}

	if err := objectRepository.SubscribeUpdateFunc(ctx, func(entities []repository.Entity) {
		for _, e := range entities {
			log.Printf("updated entity %+v", e)
		}
	}); err != nil {
		log.Fatalf("subscribe create update func failed: %s", err)
	}

	if err := objectRepository.SubscribeDeleteFunc(ctx, func(ids []repository.ID) {
		for _, id := range ids {
			log.Printf("deleted entity with id %v", id)
		}
	}); err != nil {
		log.Fatalf("subscribe create update func failed: %s", err)
	}

	if err := objectRepository.Create(ctx, []repository.Entity{
		sedmax.NewObject(1, sedmax.RootObjectID, "New object", 0),
		sedmax.NewObject(777, sedmax.RootObjectID, "New object", 0),
	}); err != nil {
		log.Fatalf("create failed: %s", err)
	}

	if err := objectRepository.Update(ctx, []repository.Entity{
		sedmax.NewObject(777, sedmax.RootObjectID, "New object", 0),
	}); err != nil {
		log.Fatalf("update failed: %s", err)
	}

	if err := objectRepository.Update(ctx, []repository.Entity{
		sedmax.NewObject(777, sedmax.RootObjectID, "My object", 0),
	}); err != nil {
		log.Fatalf("update failed: %s", err)
	}

	if err := objectRepository.Delete(ctx, []repository.ID{
		sedmax.ObjectID(777),
	}); err != nil {
		log.Fatalf("delete failed: %s", err)
	}

	entities, err := objectRepository.Read(ctx)
	if err != nil {
		log.Fatalf("read failed: %s", err)
	}
	for _, e := range entities {
		log.Printf("readed entity %v", e)
	}
}
