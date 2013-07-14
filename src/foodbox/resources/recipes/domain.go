package recipes

import (
	"foodbox/context"
	"appengine"
)

type Repository struct {
	queryEngine context.QueryEngine
}

type KeyedRecipe struct {
	Id string
	Recipe Recipe
}

type Recipe struct {
	Title string
	Author string
	Ingredients []Ingredient
	ImageKeys []string
}

type Ingredient struct {
	Name string
	Quantity string
	Unit string
}

const RecipeType = "Recipe"

func newRepository(queryEngine context.QueryEngine) *Repository {
	
	return &Repository{
		queryEngine,
	}
}

func newRepositoryFromContext(c appengine.Context) *Repository {
	queryEngine := context.NewQueryEngine(c)
	return newRepository(queryEngine)
}

func (repository *Repository) add(recipe *Recipe) (key string, err error) {
	key, err = repository.queryEngine.NewEntity(RecipeType, recipe)
	return;
}

func (repository *Repository) save(key string, recipe *Recipe) error {
	return repository.queryEngine.SaveEntity(RecipeType, key, recipe)
}

func (repository *Repository) fetchRecipe(recipeId string) KeyedRecipe {
	var recipe Recipe
	repository.queryEngine.GetEntity(recipeId, &recipe)
	return KeyedRecipe {
		Id: recipeId,
		Recipe: recipe,
	}
}

func (repository *Repository) removeRecipe(recipeId string) {
	repository.queryEngine.DeleteEntity(recipeId)
}

func (repository *Repository) fetchRecipes() []KeyedRecipe {
    recipes := []KeyedRecipe{}
    for t := repository.runQuery(); ; {
        var recipe Recipe
        key, more := t.Next(&recipe)

		if !more {
			break
		}
		
		model := KeyedRecipe {
			Id: key,
			Recipe: recipe,
		}
        recipes = append(recipes, model)
    }
    return recipes
}

func (repository *Repository) runQuery() context.RecordIterator {
	return repository.queryEngine.NewQuery(RecipeType)
}
