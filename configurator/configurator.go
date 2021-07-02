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

type CreateSubscriber interface {
	SubscribeCreateFunc(context.Context, func(entities []Entity)) error
}

type UpdateSubscriber interface {
	SubscribeUpdateFunc(context.Context, func(entities []Entity)) error
}

type DeleteSubscriber interface {
	SubscribeDeleteFunc(context.Context, func(ids []ID)) error
}
