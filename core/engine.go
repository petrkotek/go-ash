package core

import (
	"errors"
	"reflect"
	"time"
)

/**
 * The heart of the Entity framework. It is responsible for keeping track of Entity objects and
 * managing EntitySystem objects. The Engine should be updated every tick via the Update(float) method.
 *
 * With the Engine you can:
 *
 * - Add/Remove Entity objects
 * - Add/Remove EntitySystem objects
 * - Obtain a list of entities for a specific Family
 * - Update the main loop
 * - Register/unregister EntityListener objects
 */
type Engine struct {
	entityManager *EntityManager
	systemManager *SystemManager
	familyManager *FamilyManager
	updating      bool
}

var ErrUpdating = errors.New("Cannot call Update() on an Engine that is already updating.")

func NewEngine() *Engine {
	engine := &Engine{}
	entityManager := NewEntityManager(engine)
	engine.entityManager = entityManager
	engine.systemManager = NewSystemManager()
	engine.familyManager = NewFamilyManager(entityManager.Entities())
	return engine
}

func (e *Engine) AddEntity(entity *Entity) {
	delayed := e.updating || e.familyManager.Notifying()
	e.entityManager.AddEntity(entity, delayed)
}

func (e *Engine) RemoveEntity(entity *Entity) {
	e.entityManager.RemoveEntity(entity, e.updating)
}

func (e *Engine) EntitiesFor(family *Family) []*Entity {
	return e.familyManager.EntitiesFor(family)
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

func (e *Engine) OnEntityAdded(entity *Entity) {
	e.addEntityInternal(entity)
}

func (e *Engine) addEntityInternal(entity *Entity) {
	e.familyManager.UpdateFamilyMembership(entity)
}

func (e *Engine) OnEntityRemoved(entity *Entity) {
	// TODO
}
