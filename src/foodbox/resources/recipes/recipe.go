package recipes

import (
	"net/http"
	"github.com/gorilla/mux"
	"foodbox/response"
	"foodbox/context"
	"appengine"
)

func getRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeId := getRecipeId(r)
	
	c := appengine.NewContext(r)
	service := NewService(c)

	recipe := service.FetchRecipe(recipeId)
	response.WriteJson(w, recipe)
}

func deleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeId := getRecipeId(r)
	
	c := appengine.NewContext(r)
	service := NewService(c)
	
	service.RemoveRecipe(recipeId)
}

type Image struct {
	Url ImageUrl
}

func serveImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blobKey := vars["id"]
	
	c := appengine.NewContext(r)
	fileStore := context.NewFileStore(c)
	fileStore.WriteFile(w, blobKey)
}

func addImageHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	fileStore := context.NewFileStore(c)
	_, imageKey := fileStore.GetUploadedFileKey(r, "file")
	
	service := NewService(c)
	
	recipeId := getRecipeId(r)
	service.AddImageToRecipe(recipeId, imageKey)
	
	img := Image { serveImageRoute(imageKey) }
	response.WriteJson(w, img)
}

func getRecipeId(r *http.Request) string {
	vars := mux.Vars(r)
	recipeId := vars["recipeId"]
	return recipeId
}