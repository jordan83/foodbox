package bulk

import (
	"net/http"
	"appengine"
	"foodbox/context"
	"foodbox/response"
	"foodbox/resources/recipes"
)

func bulkUploadRecipesHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	fileStore := context.NewFileStore(c)
	
	_, key := fileStore.GetUploadedFileKey(r, "file")
	reader := fileStore.GetReader(key)
	
	var uploadedRecipes []recipes.Recipe
	response.DecodeJson(reader, &uploadedRecipes)
	
	service := recipes.NewService(c)
	var models []recipes.RecipeModel
	for _, recipe := range uploadedRecipes {
		model, _ := service.Add(&recipe)
		models = append(models, model) 
	}
	
	fileStore.RemoveFiles([]string{key}) 
	
	response.WriteJson(w, models)
}


func bulkUploadInitHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	fileStore := context.NewFileStore(c)
	
	w.Write([]byte(fileStore.CreateUploadUrl(uploadStartRoute)))
}