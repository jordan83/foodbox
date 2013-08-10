package bulk

import (
	"github.com/gorilla/mux"
)

func InitRoutes(router *mux.Router) {
	router.HandleFunc(uploadStartRoute, bulkUploadRecipesHandler).Methods("POST")
	router.HandleFunc("/upload/init", bulkUploadInitHandler).Methods("GET")
}

const uploadStartRoute = "/upload/start"