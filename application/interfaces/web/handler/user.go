package handler

import (
	"html/template"
	"net/http"

	"github.com/sample/application/interfaces/utils"
	"github.com/sample/application/usecase"
)

// UserHandler - user handler
type UserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler - new user handler
func NewUserHandler(userUsecase usecase.UserUsecase) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

func (h *userHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := h.userUsecase.Find(ctx)
	if err != nil {
		utils.RenderInternalServerError(w, err)
		return
	}

	tpl, err := template.ParseFiles(path+"base.html", path+"user.html")
	if err != nil {
		utils.RenderInternalServerError(w, err)
		return
	}
	err = tpl.ExecuteTemplate(w, "base", map[string]interface{}{
		"Users": users,
	})
	if err != nil {
		utils.RenderInternalServerError(w, err)
		return
	}
	return
}
