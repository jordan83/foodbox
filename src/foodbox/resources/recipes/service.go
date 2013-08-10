package recipes

import (
	"foodbox/context"
	"appengine"
)

//----------------------------------------------------------------------
// Datastructures and private methods for default service implementation
//----------------------------------------------------------------------

type RepoWrapper struct {
	fileStore context.FileStore
	repository *Repository
}

type RecipeModel struct {
	Id string
	Title string
	Author string
	FileUploadUrl string
	Ingredients []Ingredient
	ImageUrls []ImageUrl
}

type ImageUrl string

func NewService(c appengine.Context) Service {
	return &RepoWrapper {
		context.NewFileStore(c),
		newRepositoryFromContext(c),
	}
}

func (service *RepoWrapper) newRecipeModel(key string, recipe Recipe) RecipeModel {
	return RecipeModel {
		Id: key,
		Title: recipe.Title,
		Author: recipe.Author,
		Ingredients: recipe.Ingredients,
		FileUploadUrl: service.fileStore.CreateUploadUrl(addFileRoute(key)),
		ImageUrls: transformToUrls(recipe.ImageKeys),
	}
}

func (service *RepoWrapper) transformToRecipeModels(keyedRecipes []KeyedRecipe) []RecipeModel {
	recipes := []RecipeModel{}
	for i := 0; i < len(keyedRecipes); i++ {
		keyedRecipe := keyedRecipes[i]
		recipes = append(recipes, service.newRecipeModel(keyedRecipe.Id, keyedRecipe.Recipe))
	}
	return recipes
}

func transformToUrls(keys []string) []ImageUrl {
	imageUrls := []ImageUrl{}
	for i := 0; i < len(keys); i++ {
		imageUrls = append(imageUrls, serveImageRoute(keys[i]))
	}
	return imageUrls
}

//------------------------------
// Recipe service implementation
//------------------------------

type Service interface {
	
	FetchRecipe(recipeId string) RecipeModel
	
	FetchRecipes() []RecipeModel
	
	Add(recipe *Recipe) (r RecipeModel, err error)
	
	RemoveRecipe(recipeId string)
	
	AddImageToRecipe(recipeId string, imageKey string)
}

func (service *RepoWrapper) FetchRecipe(recipeId string) RecipeModel {
	recipe := service.repository.fetchRecipe(recipeId)
	return service.newRecipeModel(recipe.Id, recipe.Recipe)
}

func (service *RepoWrapper) FetchRecipes() []RecipeModel {
	recipes := service.repository.fetchRecipes()
	return service.transformToRecipeModels(recipes)
}

func (service *RepoWrapper) Add(recipe *Recipe) (r RecipeModel, err error) {
	key, err := service.repository.add(recipe)
	r = service.newRecipeModel(key, *recipe)
	return;
}

func (service *RepoWrapper) RemoveRecipe(recipeId string) {
	imageKeys := service.repository.fetchRecipe(recipeId).Recipe.ImageKeys
	service.fileStore.RemoveFiles(imageKeys)
	
	service.repository.removeRecipe(recipeId)
}

func (service *RepoWrapper) AddImageToRecipe(recipeId string, imageKey string) {
	keyedRecipe := service.repository.fetchRecipe(recipeId)
	
	keyedRecipe.Recipe.ImageKeys = append(keyedRecipe.Recipe.ImageKeys, imageKey)
	service.repository.save(keyedRecipe.Id, &keyedRecipe.Recipe)
}

