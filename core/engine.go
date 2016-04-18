package core

type Engine struct {
	entityManager *EntityManager
	updating bool
}

func NewEngine() *Engine {
	return &Engine{
		entityManager: NewEntityManager(),
	}
}

func (e *Engine) AddEntity(entity *Entity) {
	delayed := e.updating // || e.familyManager.Notifying()
	e.entityManager.AddEntity(entity, delayed)
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
