package http_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/perebaj/credit"
	"github.com/perebaj/credit/bureaus"
	"github.com/perebaj/credit/http"
	"github.com/perebaj/credit/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestHandlerSaveCompany(t *testing.T) {
	ctrl := gomock.NewController(t)
	companySvcMock := mock.NewMockCompanyService(ctrl)
	companySvcMock.EXPECT().SaveCompany(gomock.Any(), gomock.Any()).Return(nil)

	bureauSvcMock := mock.NewMockBureauService(ctrl)
	bureauSvcMock.EXPECT().Fetch(gomock.Any(), gomock.Any()).Return(bureaus.Empresa{
		NomeEmpresarial: "Company Jojo",
	}, nil)

	handler := http.NewHandler(companySvcMock, bureauSvcMock)
	mux := handler.Router()

	req := httptest.NewRequest("GET", "/company?cnpj=123&cpf=321", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	require.Equal(t, 201, w.Code)
}

func TestHandlerSaveCompany_invalidMethod(t *testing.T) {
	ctrl := gomock.NewController(t)
	companySvcMock := mock.NewMockCompanyService(ctrl)

	handler := http.NewHandler(companySvcMock, nil)
	mux := handler.Router()

	payload := credit.Company{
		Name: "Company",
		ID:   "123",
	}

	//convert struct to io.Reader
	body, err := json.Marshal(payload)
	require.NoError(t, err)
	buf := bytes.NewBuffer(body)

	req := httptest.NewRequest("POST", "/company", buf)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)

	require.Equal(t, 405, w.Code)
}
