package core

import "gopkg.in/fatih/set.v0"

type FamilyManager struct {
	entitySet         *set.Set
	pendingOperations []*EntityOperation
	immutableFamilies map[*Family][]*Entity
	notifying         bool
}

func NewFamilyManager(entities *set.Set) *FamilyManager {
	return &FamilyManager{
		entitySet:         entities,
		immutableFamilies: make(map[*Family][]*Entity),
	}
}

func (fm *FamilyManager) Notifying() bool {
	return fm.notifying
}

func (fm *FamilyManager) EntitiesFor(family *Family) []*Entity {
	return fm.registerFamily(family)
}

func (fm *FamilyManager) registerFamily(family *Family) []*Entity {
	entitiesInFamily, ok := fm.immutableFamilies[family]
	if !ok {
		entitiesInFamily = make([]*Entity, 0)
		fm.entitySet.Each(func(ent interface{}) bool {
			entity := ent.(*Entity)
			if family.Matches(entity) {
				entitiesInFamily = append(entitiesInFamily, entity)
			}
			return true
		})
		fm.immutableFamilies[family] = entitiesInFamily
	}
	return entitiesInFamily
}

func (fm *FamilyManager) UpdateFamilyMembership(entity *Entity, added bool) {
	for family, entities := range fm.immutableFamilies {
		if family.Matches(entity) {
			if added {
				fm.immutableFamilies[family] = append(entities, entity)
			} else {
				slice := fm.immutableFamilies[family]
				for pos, e := range slice {
					if e == entity {
						slice = append(slice[:pos], slice[pos+1:]...)
						fm.immutableFamilies[family] = slice
						break
					}
				}
			}
		}
	}
}
