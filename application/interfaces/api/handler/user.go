package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sample/application/entity/model"
	"github.com/sample/application/interfaces/utils"
	"github.com/sample/application/usecase"
)

// UserHandler - user handler
type UserHandler interface {
	Get(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
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
	userID, err := strconv.ParseUint(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		utils.RenderBadRequest(w)
		return
	}
	user, err := h.userUsecase.Get(ctx, userID)
	if err != nil {
		utils.RenderInternalServerError(w, err)
		return
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("json marshal error: %+v\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	return
}

func (h *userHandler) Find(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	users, err := h.userUsecase.Find(ctx)
	if err != nil {
		utils.RenderInternalServerError(w, err)
		return
	}
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("json marshal error: %+v\n", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
	return
}

// UserCreateParam - user create param
type UserCreateParam struct {
	Name string `json:"name"`
}

func (h *userHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Body == nil {
		utils.RenderBadRequest(w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	param := UserCreateParam{}
	err := decoder.Decode(&param)
	if err != nil {
		log.Println("body decode failed")
		return
	}
	user := &model.User{
		Name: param.Name,
	}

	user, err = h.userUsecase.Create(ctx, user)
	if err != nil {
		utils.RenderInternalServerError(w, err)
		return
	}
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("json marshal error: %+v\n", err)
		utils.RenderInternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// UserUpdateParam - user upadte param
type UserUpdateParam struct {
	Name string `json:"name"`
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.ParseUint(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		utils.RenderBadRequest(w)
		return
	}

	if r.Body == nil {
		utils.RenderBadRequest(w)
		return
	}
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	param := UserUpdateParam{}
	err = decoder.Decode(&param)
	if err != nil {
		log.Println("body decode failed")
		utils.RenderInternalServerError(w, err)
		return
	}
	user := &model.User{
		ID:   userID,
		Name: param.Name,
	}

	err = h.userUsecase.Update(ctx, user)
	if err != nil {
		utils.RenderInternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *userHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := strconv.ParseUint(chi.URLParam(r, "userID"), 10, 64)
	if err != nil {
		utils.RenderBadRequest(w)
		return
	}

	err = h.userUsecase.Delete(ctx, userID)
	if err != nil {
		utils.RenderInternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
