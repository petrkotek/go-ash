package core

import (
	"errors"

	"gopkg.in/fatih/set.v0"
)

var ErrAlreadyRegistered = errors.New("Entity is already registered")

type EntityManager struct {
	pendingOperations []*EntityOperation
	entitySet         *set.Set
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		entitySet: set.New(),
	}
}

func (em *EntityManager) AddEntity(entity *Entity, delayed bool) {
	if delayed {
		operation := &EntityOperation{Add, entity}
		em.pendingOperations = append(em.pendingOperations, operation);
		return
	}
	em.addEntityInternal(entity);
}

func (em *EntityManager) Entities() *set.Set {
	return em.entitySet
}

func (em *EntityManager) addEntityInternal(entity *Entity) error {
	if (em.entitySet.Has(entity)) {
		return ErrAlreadyRegistered
	}

	em.entitySet.Add(entity);

	// TODO: listener.entityAdded(entity);
	return nil
}

type EntityOperationType int

const (
	Add EntityOperationType = iota
	Remove
	// TODO: RemoveAll
)

type EntityOperation struct {
	entityOperationType EntityOperationType
	entity              *Entity
}
