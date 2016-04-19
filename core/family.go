package core

import (
	"gopkg.in/fatih/set.v0"
)

type Family struct {
	components *set.Set
}

func NewFamily(components ...interface{}) *Family {
	return &Family{
		components: set.New(components...),
	}
}

func (f *Family) Matches(entity *Entity) bool {
	for _, requiredComponent := range f.components.List() {
		if entity.Component(requiredComponent) == nil {
			return false
		}
	}
	return true
}
