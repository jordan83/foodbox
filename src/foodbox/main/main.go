package main

import (
    "net/http"
    "foodbox/response"
    "foodbox/resources/recipes"
    "foodbox/context"
    "github.com/gorilla/mux"
)

type IndexData struct{
	Title string
	UserUrl string
	UserName string
}

func init() {
	r := mux.NewRouter()
    r.HandleFunc("/", mainHandler)
    r.HandleFunc("/partial/recipes", recipesHandler)
    r.HandleFunc("/partial/recipe", recipeHandler)
    r.HandleFunc("/partial/createRecipe", createRecipesHandler)
    r.HandleFunc("/nav", navHandler)
    
    recipes.InitRoutes(r)
    
    http.Handle("/", r)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {	
	foodboxContext := context.NewContext(r)
	redirected := foodboxContext.RedirectIfNotLoggedIn(w)
	if redirected {
		return
	}
	
	url := foodboxContext.LogoutURL()
	d := IndexData {
		Title: "Foodbox",
		UserUrl: url,
		UserName: foodboxContext.CurrentUser.String(),
	}
	response.RenderTemplate("index.html", w, d)
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	response.RenderHtml("recipes.html", w)
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
	response.RenderHtml("recipe.html", w)
}

type Nav struct {
	Route string
	Name string
}

func navHandler(w http.ResponseWriter, r *http.Request) {
	items := [1]Nav{ Nav{ Route: "/", Name: "Home" }}
	response.WriteJson(w, items)
}

func createRecipesHandler(w http.ResponseWriter, r *http.Request) {
	response.RenderHtml("createRecipe.html", w)
}