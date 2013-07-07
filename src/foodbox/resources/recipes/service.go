package recipes

import (
	"foodbox/context"
	"appengine"
)

// Repository wrapper implementation.

type RepoWrapper struct {
	fileUploadUrlProvider context.FileUploadUrlProvider
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

func newService(c appengine.Context) Service {
	return &RepoWrapper {
		context.NewFileUploadUrlProvider(c),
		newRepositoryFromContext(c),
	}
}

func (service *RepoWrapper) newRecipeModel(key string, recipe Recipe) RecipeModel {
	return RecipeModel {
		Id: key,
		Title: recipe.Title,
		Author: recipe.Author,
		Ingredients: recipe.Ingredients,
		FileUploadUrl: service.fileUploadUrlProvider.CreateUploadUrl(addFileRoute(key)),
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

// Recipe service implementation

type Service interface {
	
	fetchRecipe(recipeId string) RecipeModel
	
	fetchRecipes() []RecipeModel
	
	add(recipe *Recipe) error
	
	removeRecipe(recipeId string)
	
	addImageToRecipe(recipeId string, imageKey string)
}

func (service *RepoWrapper) fetchRecipe(recipeId string) RecipeModel {
	recipe := service.repository.fetchRecipe(recipeId)
	return service.newRecipeModel(recipe.Id, recipe.Recipe)
}

func (service *RepoWrapper) fetchRecipes() []RecipeModel {
	recipes := service.repository.fetchRecipes()
	return service.transformToRecipeModels(recipes)
}

func (service *RepoWrapper) add(recipe *Recipe) error {
	return service.repository.add(recipe)
}

func (service *RepoWrapper) removeRecipe(recipeId string) {
	service.repository.removeRecipe(recipeId)
}

func (service *RepoWrapper) addImageToRecipe(recipeId string, imageKey string) {
	keyedRecipe := service.repository.fetchRecipe(recipeId)
	
	keyedRecipe.Recipe.ImageKeys = append(keyedRecipe.Recipe.ImageKeys, imageKey)
	service.repository.save(keyedRecipe.Id, &keyedRecipe.Recipe)
}

