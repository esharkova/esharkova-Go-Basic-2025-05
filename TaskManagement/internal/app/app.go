package app

import (
	"TaskManagement/internal/handler"
	"TaskManagement/internal/usecase"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"log"
)

type App struct {
	router *gin.Engine
	server *http.Server
}

func New() *App {
	router := gin.Default()

	uc := usecase.New()

	h := handler.New(uc)

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/tasks", h.GetTasks())
	v1.GET("/task/:id", h.GetTask())
	v1.POST("/task", h.CreateTask())
	v1.PUT("/task/:id", h.UpdateTask())
	v1.DELETE("/task/:id", h.DeleteTask())

	v1.GET("/users", h.GetUsers())
	v1.GET("/user/:id", h.GetUser())
	v1.POST("/user", h.CreateUser())
	v1.PUT("/user/:id", h.UpdateUser())
	v1.DELETE("/user/:id", h.DeleteUser())

	return &App{
		router: router,
		server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
	}
}

func (a *App) Start() {
	go func() {
		log.Printf("Server starting on %s", ":8080")
		if err := a.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()
}

func (a *App) Stop() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
