package main

import (
	"context"
	"log"

	"github.com/rekhin/generator/inmemory"
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
)

func main() {
	ctx := context.TODO()

	// // TODO реализовать в следующем виде! см. ниже
	// // start
	// var (
	// 	inmemoryRepo = inmemory.NewRepository()
	// 	// inmemoryPub  = inmemory.NewPublisher(inmemoryRepo)
	// 	inmemorySub = inmemory.NewSubscriber(inmemoryRepo)

	// 	grpcRepo = grpc.NewRepository()
	// 	grpcPub  = grpc.NewPublisher(grpcRepo)
	// 	// grpcSub  = grpc.NewSubscriber(grpcRepo)
	// )

	// inmemorySub.SubscribeDeltaFunc(ctx, grpcPub.PublishDelta)
	// inmemorySub.SubscribeDeltaHandler(ctx, grpcPub)
	// // end

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
		sedmax.NewObject(1, sedmax.RootObjectID, "New object", 0),
		sedmax.NewObject(777, sedmax.RootObjectID, "New object", 0),
	); err != nil {
		log.Fatalf("create failed: %s", err)
	}

	if err := objectInmemoryRepository.UpdateEntities(ctx, sedmax.NewObject(777, sedmax.RootObjectID, "New object", 0)); err != nil {
		log.Fatalf("update failed: %s", err)
	}

	if err := objectInmemoryRepository.UpdateEntities(ctx, sedmax.NewObject(777, sedmax.RootObjectID, "My object", 0)); err != nil {
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
		log.Printf("read entity %v", e)
	}
}
