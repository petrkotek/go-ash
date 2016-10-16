package core

import (
	"errors"
	"reflect"
	"time"
)

type Engine struct {
	entityManager *EntityManager
	systemManager *SystemManager
	updating      bool
}

var ErrUpdating = errors.New("Cannot call Update() on an Engine that is already updating.")

func NewEngine() *Engine {
	return &Engine{
		entityManager: NewEntityManager(),
		systemManager: NewSystemManager(),
	}
}

func (e *Engine) AddEntity(entity *Entity) {
	delayed := e.updating
	e.entityManager.AddEntity(entity, delayed)
}

func (e *Engine) RemoveEntity(entity *Entity) {
	e.entityManager.RemoveEntity(entity, e.updating)
}

func (e *Engine) EntitiesFor(family *Family) []*Entity {
	result := []*Entity{}
	entities := e.entityManager.Entities()
	entities.Each(func(ent interface{}) bool {
		entity := ent.(*Entity)
		if family.Matches(entity) {
			result = append(result, entity)
		}
		return true
	})
	return result
}

func (e *Engine) AddSystem(system EntitySystem) {
	e.systemManager.Add(system)
}

func (e *Engine) System(system EntitySystem) EntitySystem {
	for _, s := range e.systemManager.Systems() {
		if reflect.TypeOf(s) == reflect.TypeOf(system) {
			return s
		}
	}
	return nil
}

func (e *Engine) Update(deltaTime time.Duration) error {
	if e.updating {
		return ErrUpdating
	}

	e.updating = true
	for _, system := range e.systemManager.Systems() {
		if system.CheckProcessing() {
			system.Update(float32(deltaTime) / float32(time.Second))
		}

		e.entityManager.ProcessPendingOperations()
	}

	e.updating = false
	return nil
}
