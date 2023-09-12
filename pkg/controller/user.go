package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ryohei1216/go-mvc/pkg/model"
	"github.com/ryohei1216/go-mvc/pkg/view"
	"gorm.io/gorm"
)

type UserController interface {
	Get(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type userControllerImpl struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) UserController {
	return &userControllerImpl{
		db: db,
	}
}

func (u *userControllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var user model.User
	if err := u.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			res := view.ErrorResponse{
				Message: "not found",
			}

			handleErrorResponse(w, res, http.StatusNotFound)
			return
		}

		res := view.ErrorResponse{
			Message: "internal server error",
		}

		handleErrorResponse(w, res, http.StatusInternalServerError)
		return
	}

	viewUser := view.User{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	handleResponse(w, viewUser, http.StatusOK)
}

func (u *userControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var req view.CreateUserRequest

	dec := json.NewDecoder(r.Body)
	dec.UseNumber()

	if err := dec.Decode(&req); err != nil {
		res := view.ErrorResponse{
			Message: "bad request",
		}

		handleErrorResponse(w, res, http.StatusBadRequest)
		return
	}

	user := model.User{
		Name: req.Name,
		Age:  req.Age,
	}

	if err := u.db.Create(&user).Error; err != nil {
		res := view.ErrorResponse{
			Message: "internal server error",
		}

		handleErrorResponse(w, res, http.StatusInternalServerError)
		return
	}

	viewUser := view.User{
		ID:        user.ID,
		Name:      user.Name,
		Age:       user.Age,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	handleResponse(w, viewUser, http.StatusOK)
}

func handleResponse(w http.ResponseWriter, res view.User, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.Warn(fmt.Sprintf("encode response: %v", err))
	}
}

func handleErrorResponse(w http.ResponseWriter, res view.ErrorResponse, status int) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		slog.Warn(fmt.Sprintf("encode response: %v", err))
	}
}
