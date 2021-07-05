package main

import (
	"context"
	"log"

	"github.com/rekhin/generator/inmemory"
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
	"github.com/rekhin/generator/types"
)

func main() {
	ctx := context.TODO()

	objectInmemoryRepository := inmemory.NewRepository()
	// objectGrpcRepository := inmemory.NewRepository()

	if err := objectInmemoryRepository.SubscribeDelta(ctx, func(d repository.Delta) {
		for _, createEntity := range d.CreateEntities {
			log.Printf("create entity %+v", createEntity)
		}
		for _, updateEntity := range d.UpdateEntities {
			log.Printf("update entity %+v", updateEntity)
		}
		for _, deleteEntityWithID := range d.DeleteEntitiesWithIDs {
			log.Printf("delete entity with id %v", deleteEntityWithID)
		}

		// if err := objectGrpcRepository.PublishDelta(ctx, d); err != nil {
		// 	log.Fatalf("publish entities failed: %s", err)
		// }
	}); err != nil {
		log.Fatalf("subscribe create update func failed: %s", err)
	}

	if err := objectInmemoryRepository.CreateEntities(ctx,
		types.Object{
			Node: types.Node{
				ID:       sedmax.ObjectID(1),
				ParentID: sedmax.ObjectID(sedmax.RootObjectID),
				Name:     "New object",
				Sort:     0,
			},
		},
		types.Object{
			Node: types.Node{
				ID:       sedmax.ObjectID(777),
				ParentID: sedmax.ObjectID(sedmax.RootObjectID),
				Name:     "New object",
				Sort:     1,
			},
		},
	); err != nil {
		log.Fatalf("create failed: %s", err)
	}

	if err := objectInmemoryRepository.UpdateEntities(ctx,
		types.Object{
			Node: types.Node{
				ID:       sedmax.ObjectID(777),
				ParentID: sedmax.ObjectID(sedmax.RootObjectID),
				Name:     "New object",
				Sort:     1,
			},
		}); err != nil {
		log.Fatalf("update failed: %s", err)
	}

	if err := objectInmemoryRepository.UpdateEntities(ctx,
		types.Object{
			Node: types.Node{
				ID:       sedmax.ObjectID(777),
				ParentID: sedmax.ObjectID(sedmax.RootObjectID),
				Name:     "My object",
				Sort:     1,
			},
		}); err != nil {
		log.Fatalf("update failed: %s", err)
	}

	if err := objectInmemoryRepository.DeleteEntitiesWithIDs(ctx, sedmax.ObjectID(777)); err != nil {
		log.Fatalf("delete failed: %s", err)
	}

	entities, err := objectInmemoryRepository.ReadEntities(ctx)
	if err != nil {
		log.Fatalf("read failed: %s", err)
	}
	for _, e := range entities {
		log.Printf("read entity %+v", e)
	}
}
