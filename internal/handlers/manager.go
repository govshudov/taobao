package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"taobao/internal/handlers/http"
	"taobao/internal/services"
)

const (
	baseURL     = "/api"
	productsURL = baseURL + "/open/product"
)

func Manager(ctx context.Context, taobaoUrl string) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route(productsURL, func(subRouter chi.Router) {
		productService := services.NewProductService(taobaoUrl)
		productHandler := http.NewProductHandler(ctx, productService)
		productHandler.ProductRegisterRoutes(subRouter)
	})

	return r
}
