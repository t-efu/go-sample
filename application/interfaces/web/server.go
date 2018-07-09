package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sample/application/repository"
	"github.com/sample/application/usecase"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sample/application/interfaces/web/handler"
)

func main() {
	endpoint := os.Getenv("DB_CONNECT_INFO")
	db, err := gorm.Open("mysql", endpoint)
	if err != nil {
		log.Panic("failed database connect.err:", err)
	}
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)
	indexHandler := handler.NewIndexHandler()

	r := chi.NewRouter()
	r.Get("/", indexHandler.Get)
	r.Get("/users", userHandler.Get)
	http.ListenAndServe(":8888", r)
}
