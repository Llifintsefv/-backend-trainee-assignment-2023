package api

import (
	"backend-trainee-assignment-2023/internal/core/interfaces"
	"backend-trainee-assignment-2023/internal/core/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	segmentId, err := h.segmentService.CreateSegment(ctx, segment)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create segment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	
	response := models.SegmentResponse{
		Status: "success",
		Id: segmentId,
		Data: models.Segment{
			Slug: segment.Slug,
			AutoAddPercent: segment.AutoAddPercent,
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteSegment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	ctx := r.Context()

	err := h.segmentService.DeleteSegment(ctx,slug)
	if err != nil {
		http.Error(w, "Failed to delete segment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

	response := models.UserResponse{
		Status: "success",
		Id: user.Id,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}


func (h *Handler) CreateSegmentUser(w http.ResponseWriter, r *http.Request){
	var req models.UserSegmentRequest
	vars := mux.Vars(r)
	userId,err := strconv.ParseInt(vars["user_id"],10,64)
	if err != nil {
		http.Error(w, "Failed to parse user_id", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Failed to decode request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	err = h.segmentService.CreateUserSegment(ctx,int(userId),req.Add,req.Remove,req.TTL)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to create user segment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetUserSegments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId,err := strconv.ParseInt(vars["user_id"],10,64)
	if err != nil {
		http.Error(w, "Failed to parse user_id", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	
	segments, err := h.segmentService.GetUserSegments(ctx,int(userId))
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get user segments", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(segments); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) DeleteUserSegment(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId,err := strconv.ParseInt(vars["user_id"],10,64)
	if err != nil {
		http.Error(w, "Failed to parse user_id", http.StatusBadRequest)
		return
	}
	segmentId,err := strconv.ParseInt(vars["segment_id"],10,64)
	if err != nil {
		http.Error(w, "Failed to parse segment_id", http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	err = h.segmentService.DeleteUserSegment(ctx,int(userId),int(segmentId))
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete user segment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)	
}

