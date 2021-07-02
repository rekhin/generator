package sedmax

import "github.com/rekhin/generator/configurator"

type Object interface {
	Node
}

type ObjectID int

const RootObjectID = 0

func (id ObjectID) Category() Category {
	return CategoryObject
}

type object struct {
	id       ObjectID
	parentID ObjectID
	name     string
	sort     int
}

func NewObject(id, parentID ObjectID, name string, sort int) Object {
	return &object{
		id:       id,
		parentID: parentID,
		name:     name,
		sort:     sort,
	}
}

func (o *object) Equal(e configurator.Entity) bool {
	o2, ok := e.(*object)
	if !ok {
		return false
	}
	if o2.id != o.id {
		return false
	}
	if o2.parentID != o.parentID {
		return false
	}
	if o2.name != o.name {
		return false
	}
	if o2.sort != o.sort {
		return false
	}
	return true
}

func (o *object) ID() configurator.ID {
	return o.id
}

func (o *object) ParentID() configurator.ID {
	return o.parentID
}

func (o *object) Name() string {
	return o.name
}

func (o *object) Sort() int {
	return o.sort
}
