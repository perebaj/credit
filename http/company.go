// Package http gather all the http handlers and routers for the credit service.
package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/perebaj/credit"
)

//go:generate mockgen -source http.go -destination ../mock/http_mock.go -package mock

// CompanyService is the interface that wraps the firestore storage methods.
type CompanyService interface {
	SaveCompany(ctx context.Context, company credit.Company) error
}

// Handler is the struct that gather all the implementions required for the http handlers.
type Handler struct {
	CompanyService CompanyService
}

func (h *Handler) saveCompany(w http.ResponseWriter, r *http.Request) {
	var company credit.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.CompanyService.SaveCompany(r.Context(), company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// NewHandler initializes a new http.Handler with the provided CompanyService.
func NewHandler(companyService CompanyService) *Handler {
	return &Handler{CompanyService: companyService}
}

// Router returns a new http.ServeMux with all the routes for the credit service.
func (h *Handler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /company", h.saveCompany)
	return mux
}
