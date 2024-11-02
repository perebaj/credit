// Package main is the entry point for the credit application.
package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/perebaj/credit/bureaus"
)

func main() {
	rfClient := bureaus.NewRFClient(&http.Client{
		Timeout: time.Second * 30,
	},
		"test-token",
	)

	rfClient.URL = "https://h-apigateway.conectagov.estaleiro.serpro.gov.br"

	_, err := rfClient.Fetch("57.348.459/0001-06", "471.092.028-10")
	if err != nil {
		slog.Info("error fetch", "error", err)
	}
}
