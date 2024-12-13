package api

import "github.com/gorilla/mux"

func SetupRouter(handler *Handler) *mux.Router {
	router := mux.NewRouter()


	router.HandleFunc("/segment/", handler.CreateSegment).Methods("POST")
	//router.HandleFunc("/api/segment/{slug}", handler.DeleteSegment).Methods("DELETE")


	return router
}