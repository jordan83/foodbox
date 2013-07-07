package recipes

import (
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
    router.HandleFunc("/recipes", getRecipesHandler).Methods("GET")
    router.HandleFunc("/recipes", createRecipeHandler).Methods("POST")
    router.HandleFunc("/recipes/{recipeId}", getRecipeHandler).Methods("GET")
    router.HandleFunc("/recipes/{recipeId}", deleteRecipeHandler).Methods("DELETE")
    router.HandleFunc("/recipes/addfile/{recipeId}", addImageHandler).Methods("POST")
    router.HandleFunc("/recipes/image/{id}", serveImageHandler).Methods("GET")
}

func addFileRoute(key string) string {
	return "/recipes/addfile/" + key
}

func serveImageRoute(imageKey string) ImageUrl {
	return ImageUrl("/recipes/image/" + imageKey)
}