package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sample/application/constant"
	"github.com/sample/application/interfaces/api/handler"
	"github.com/sample/application/repository"
	"github.com/sample/application/usecase"
)

func main() {
	db, err := gorm.Open("mysql", constant.MysqlConnectInfo)
	if err != nil {
		log.Panicln("failed database connect.err:", err)
	}
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	r := chi.NewRouter()
	r.Get("/users/{userID}", userHandler.Get)
	r.Get("/users", userHandler.Find)
	r.Post("/users", userHandler.Create)
	r.Put("/users/{userID}", userHandler.Update)
	r.Delete("/users/{userID}", userHandler.Delete)
	http.ListenAndServe(":3000", r)
}
