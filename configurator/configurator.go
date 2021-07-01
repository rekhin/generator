package configurator

import "context"

type Reader interface {
	Read(context.Context, []Entity) error
}

type Creator interface {
	Create(context.Context, []Entity) error
}

type Updater interface {
	Update(context.Context, []Entity) error
}

type Deleter interface {
	Delete(context.Context, []ID) error
}

type CreateUpdateSubscriber interface {
	SubscribeCreateUpdate(context.Context, func(entities []Entity)) error
}

type DeleteSubscriber interface {
	SubscribeDelete(context.Context, func(ids []ID)) error
}
