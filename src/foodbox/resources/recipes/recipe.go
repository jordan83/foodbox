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
	service := newService(c)

	recipe := service.fetchRecipe(recipeId)
	response.WriteJson(w, recipe)
}

func deleteRecipeHandler(w http.ResponseWriter, r *http.Request) {
	recipeId := getRecipeId(r)
	
	c := appengine.NewContext(r)
	service := newService(c)
	
	service.removeRecipe(recipeId)
}

type Image struct {
	Url ImageUrl
}

func serveImageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blobKey := vars["id"]
	
	c := appengine.NewContext(r)
	imageStore := context.NewImageStore(c)
	imageStore.WriteImage(w, blobKey)
}

func addImageHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	imageStore := context.NewImageStore(c)
	_, imageKey := imageStore.GetUploadedImageKey(r, "file")
	
	service := newService(c)
	
	recipeId := getRecipeId(r)
	service.addImageToRecipe(recipeId, imageKey)
	
	img := Image { serveImageRoute(imageKey) }
	response.WriteJson(w, img)
}

func getRecipeId(r *http.Request) string {
	vars := mux.Vars(r)
	recipeId := vars["recipeId"]
	return recipeId
}