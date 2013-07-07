package recipes

import (
	"testing"
	"foodbox/context"
)

type mockQueryEngine struct {
	entityType string
	entityKey string
	entity interface{}
	iterator context.RecordIterator
}

func (m *mockQueryEngine) NewQuery(entityType string) context.RecordIterator {
	return m.iterator
}

func (m *mockQueryEngine) NewEntity(entityType string, entity interface{}) error {
	m.entityType = entityType
	m.entity = entity
	return nil
}


func (m *mockQueryEngine) DeleteEntity(key string) {
}
	

func (m *mockQueryEngine) GetEntity(key string, entity interface{}) {
}


func (m *mockQueryEngine) SaveEntity(entityType string, entityKey string, entity interface{}) error {
	m.entityType = entityType
	m.entityKey = entityKey
	m.entity = entity
	return nil
}


func newMockQueryEngine(iterator context.RecordIterator) *mockQueryEngine {
	return &mockQueryEngine {
		entityType: "",
		entity: nil,
		iterator: iterator,
	}
}

// Todo make this work with an interface that supports a function to get an id.
type mockRecordIterator struct {
	toReturn []interface{}
	curIndex int
}

func newMockRecordIterator(toReturn []interface{}) *mockRecordIterator {
	return &mockRecordIterator {
		toReturn: toReturn,
		curIndex: 0,
	}
}

func (m *mockRecordIterator) Next(dst interface{}) (key string, more bool) {
	if m.curIndex >= len(m.toReturn) {
		key = ""
		more = false
		return
	} else {
		key = "Any"
		more = true
		dst = m.toReturn[m.curIndex]
		m.curIndex = m.curIndex + 1
		return
	}
}

func Test_NewEntityCalledWhenCreatingRecipe(t *testing.T) {
	m := newMockQueryEngine(nil)
	
	r := newRepository(m)
	
	recipe := Recipe {
		Title: "test",
		Author: "testAuthor",
		Ingredients: make([]Ingredient, 0),
	}
	
	r.add(&recipe)
	
	assertEquals(t, RecipeType, m.entityType)
	assertEquals(t, &recipe, m.entity)
}

func Test_FetchRecipesReturnsCorrectCount(t *testing.T) {
	recipe := RecipeModel {
		Id: "SomeId",
		Title: "test",
		Author: "testAuthor",
		Ingredients: make([]Ingredient, 0),
	}

	recipes := make([]interface{}, 2)
	recipes[0] = recipe
	recipes[1] = recipe
	
	r, _ := createMockRepositoryWithRecipes(recipes)
	
	allRecipes := r.fetchRecipes()
	
	assertEquals(t, 2, len(allRecipes))
}

func Test_SaveEntityCalledWhenSavingEntity(t *testing.T) {

	recipe := &Recipe {
		Title: "Any title",
		Author: "Some Author",
		Ingredients: make([]Ingredient, 0),
		ImageKeys: make([]string, 0),
	}
	
	r, qe := createMockRepositoryWithRecipes(make([]interface{}, 0))
	
	recipeKey := "anyKey"
	r.save(recipeKey, recipe) 
	
	assertEquals(t, RecipeType, qe.entityType)
	assertEquals(t, recipe, qe.entity)
	assertEquals(t, recipeKey, qe.entityKey)
}

func assertEquals(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

func createMockRepositoryWithRecipes(recipes []interface{}) (repo *Repository, qe *mockQueryEngine) {
	i := newMockRecordIterator(recipes)
	
	qe = newMockQueryEngine(i)
	repo = newRepository(qe)
	
	return
}