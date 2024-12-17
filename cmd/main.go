package main

import (
	"backend-trainee-assignment-2023/api"
	"backend-trainee-assignment-2023/internal/config"
	"backend-trainee-assignment-2023/internal/core/segment"
	"backend-trainee-assignment-2023/internal/core/storage"
	"backend-trainee-assignment-2023/internal/core/user"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfig()

	db,err := storage.NewDB(cfg.DBConnStr)
	if err != nil {
		log.Fatal()
	}

	defer db.Close()

	userRepo := user.NewUserRepo(db)
	segmentRepo := segment.NewSegmentRepository(db)
	userService := user.NewUserService(userRepo, segmentRepo)
	segmentService := segment.NewSegmentService(segmentRepo, userRepo)

	handler := api.NewHandler(userService, segmentService)

	router := api.SetupRouter(handler)

	fmt.Println("server started at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
	
}