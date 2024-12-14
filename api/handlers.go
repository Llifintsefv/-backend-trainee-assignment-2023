package api

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
	"backend-trainee-assignment-2023/internal/core/models"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	userService  interfaces.UserService
	segmentService  interfaces.SegmentService
}

func NewHandler(userService  interfaces.UserService, segmentService  interfaces.SegmentService) *Handler {
	return &Handler{userService: userService, segmentService: segmentService}
}


func (h *Handler) CreateSegment(w http.ResponseWriter, r *http.Request) {
	var segment models.Segment
	err := json.NewDecoder(r.Body).Decode(&segment) 
	if err != nil {
		http.Error(w,"Failed to decode request", http.StatusBadRequest)
	}
	ctx := r.Context()

	segmentId,err := h.segmentService.CreateSegment(ctx,segment)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create segment", http.StatusInternalServerError)
		return
	}

	
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]int{"id": segmentId}); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
	}

}


func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user) 
	if err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = h.userService.CreateUser(ctx,user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}