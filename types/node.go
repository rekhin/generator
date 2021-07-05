package types

import (
	"github.com/rekhin/generator/repository"
	"github.com/rekhin/generator/sedmax"
)

type Node struct {
	ID       sedmax.NodeID
	ParentID sedmax.NodeID
	Name     string
	Sort     int
}

func (n Node) Equal(e repository.Entity) bool {
	second, ok := e.(sedmax.Node)
	if !ok {
		return false
	}
	if n.GetID() != second.GetID() {
		return false
	}
	if n.GetParentID() != second.GetParentID() {
		return false
	}
	if n.GetName() != second.GetName() {
		return false
	}
	if n.GetSort() != second.GetSort() {
		return false
	}
	return true
}

func (n Node) GetID() repository.ID {
	return n.ID
}

func (n Node) GetParentID() repository.ID {
	return n.ParentID
}

func (n Node) GetName() string {
	return n.Name
}

func (n Node) GetSort() int {
	return n.Sort
}

var _ sedmax.Node = &Node{}
