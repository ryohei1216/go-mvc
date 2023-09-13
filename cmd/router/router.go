package router

import (
	"github.com/go-chi/chi"
	"github.com/ryohei1216/go-mvc/pkg/controller"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	uc := controller.NewUserController(db)
	r.Get("/users", uc.Get)
	r.Post("/users", uc.Create)

	return r
}
