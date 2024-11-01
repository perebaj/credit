package bureaus

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRFClient_Fetch(t *testing.T) {
	c := NewRFClient(&http.Client{
		Timeout: time.Second * 1,
	}, "test-token")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}

		if r.URL.Path != "/api-cnpj-empresa/v2/empresa/123456789" {
			t.Errorf("expected path /api-cnpj-empresa/v2/empresa/123456789, got %s", r.URL.Path)
		}

		// validate if the headers are set correctly
		if r.Header.Get("Authorization") !=
			"Bearer test-token" {
			t.Errorf("expected Authorization header to be set")
		}

		if r.Header.Get("x-cpf-usuario") != "987654321" {
			t.Errorf("expected x-cpf-usuario header to be set")
		}

		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	// Overwrite the URL with the server's URL, to make the request to the mock server instead of the real one.
	c.URL = server.URL

	_, err := c.Fetch("123456789", "987654321")
	if err != nil {
		t.Fatalf("could not fetch data from Receita Federal: %v", err)
	}
}
