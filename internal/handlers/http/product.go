package http

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"taobao/internal/dto"
	"taobao/internal/services"
)

type ProductsHandler struct {
	ctx     context.Context
	service services.ProductRepository
}

func NewProductHandler(ctx context.Context, service services.ProductRepository) *ProductsHandler {
	return &ProductsHandler{
		ctx:     ctx,
		service: service,
	}
}

func (h *ProductsHandler) ProductRegisterRoutes(r chi.Router) {
	r.Method(http.MethodPost, "/recommend", http.HandlerFunc(h.v1RecommendedProducts))
	r.Method(http.MethodGet, "/get", http.HandlerFunc(h.v1ProductDetails))
}

func (h *ProductsHandler) v1RecommendedProducts(w http.ResponseWriter, r *http.Request) {
	var req dto.RecommendedProductsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	res, err := h.service.RecommendedProducts(r.Context(), req.PageNo, req.PageSize, req.Language)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, res, http.StatusOK)
}

func (h *ProductsHandler) v1ProductDetails(w http.ResponseWriter, r *http.Request) {
	itemId, err := strconv.ParseInt(r.URL.Query().Get("itemId"), 10, 64)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := h.service.ProductDetails(r.Context(), itemId)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	h.writeJSON(w, res, http.StatusOK)
}

func (h *ProductsHandler) writeJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *ProductsHandler) writeError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	}); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}
