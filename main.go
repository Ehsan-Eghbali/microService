package main

import (
	"log"
	"microService/config"
	"microService/controller"
	"microService/db"
	"microService/middleware"
	"microService/repository"
	"microService/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	dbConn := db.ConnectDB(cfg.DbURL)

	userRepo := repository.NewUserRepository(dbConn)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := mux.NewRouter()

	r.HandleFunc("/users", userController.GetUsers).Methods("GET")
	r.HandleFunc("/user", userController.GetUser).Methods("GET")
	r.HandleFunc("/users", userController.CreateUser).Methods("POST")

	r.Use(middleware.LoggingMiddleware)

	log.Printf("Server is running on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatalf("Could not start server: %s", err.Error())
	}
}
