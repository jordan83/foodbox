package recipes

import (
	"net/http"
	"foodbox/response"
	"appengine"
)

func getRecipesHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    service := NewService(c)

    recipes := service.FetchRecipes()    
	response.WriteJson(w, recipes)
}

func createRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	response.DecodeJson(r.Body, &recipe)
	
	c := appengine.NewContext(r)
	service := NewService(c)
	recipeModel, err := service.Add(&recipe)
	
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    response.WriteJson(w, recipeModel)
}