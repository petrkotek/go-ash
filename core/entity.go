package core

import (
	"reflect"

	"gopkg.in/fatih/set.v0"
)

type Entity struct {
	components          *set.Set
	scheduledForRemoval bool
	removing            bool
}

func NewEntity() *Entity {
	return &Entity{
		components: set.New(),
	}
}

func (e *Entity) Add(component Component) {
	e.components.Add(component)
}

func (e *Entity) Components() *set.Set {
	return e.components
}

func (e *Entity) Component(component Component) Component {
	for _, c := range e.components.List() {
		if reflect.TypeOf(c) == reflect.TypeOf(component) {
			return c
		}
	}
	return nil
}
