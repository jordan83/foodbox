package recipes

import (
	"net/http"
	"foodbox/response"
	"appengine"
)

func getRecipesHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
    service := newService(c)

    recipes := service.fetchRecipes()    
	response.WriteJson(w, recipes)
}

func createRecipeHandler(w http.ResponseWriter, r *http.Request) {
	var recipe Recipe
	response.DecodeJson(r.Body, &recipe)
	
	c := appengine.NewContext(r)
	service := newService(c)
	err := service.add(&recipe)
	
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}