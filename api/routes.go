package api

import "github.com/gorilla/mux"

func SetupRouter(handler *Handler) *mux.Router {
	router := mux.NewRouter()


	router.HandleFunc("/segment/", handler.CreateSegment).Methods("POST")
	router.HandleFunc("/segment/{slug}", handler.DeleteSegment).Methods("DELETE")
	router.HandleFunc("/user/", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users/{user_id}/segments", handler.CreateSegmentUser).Methods("POST")
	router.HandleFunc("/users/{user_id}/segments",handler.GetUserSegments).Methods("GET")
	router.HandleFunc("/segments/{slug}/users", handler.DeleteUserSegment).Methods("DELETE")



	return router
}