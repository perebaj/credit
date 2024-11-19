// Package http gather all the http handlers and routers for the credit service.
package http

import (
	"context"
	"net/http"

	"github.com/perebaj/credit"
	"github.com/perebaj/credit/bureaus"
)

//go:generate mockgen -source company.go -destination ../mock/http_mock.go -package mock

// CompanyService is the interface that wraps the firestore storage methods.
type CompanyService interface {
	SaveCompany(ctx context.Context, company credit.Company) error
}

// BureauService is the interface that wraps the methods to fetch data from the bureaus.
type BureauService interface {
	Fetch(cnpj string, cpf string) (bureaus.Empresa, error)
}

// Handler is the struct that gather all the implementions required for the http handlers.
type Handler struct {
	CompanyService CompanyService
	BureauService  BureauService
}

func (h *Handler) saveCompany(w http.ResponseWriter, r *http.Request) {
	//read cnpj and cpf from request
	cnpj := r.URL.Query().Get("cnpj")
	cpf := r.URL.Query().Get("cpf")

	empresa, err := h.BureauService.Fetch(cnpj, cpf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	company := credit.Company{
		ID:   cnpj,
		Name: empresa.NomeEmpresarial,
	}

	if err := h.CompanyService.SaveCompany(r.Context(), company); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// NewHandler initializes a new http.Handler with the provided CompanyService.
func NewHandler(companyService CompanyService, bureauService BureauService) *Handler {
	return &Handler{CompanyService: companyService, BureauService: bureauService}
}

// Router returns a new http.ServeMux with all the routes for the credit service.
func (h *Handler) Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /company", h.saveCompany)
	return mux
}
