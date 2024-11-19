// Package firestore gather all function that deal with the firestore database
package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/perebaj/credit"
)

// Storage is a struct that holds the firestore client and the projectID and database name
type Storage struct {
	client    *firestore.Client
	projectID string
}

// NewStorage creates a new Storage struct
func NewStorage(client *firestore.Client, projectID string) *Storage {
	return &Storage{
		client:    client,
		projectID: projectID,
	}
}

// SaveCompany save a company in the firestore database
func (s *Storage) SaveCompany(ctx context.Context, company credit.Company) error {
	collection := s.client.Collection("companies")
	_, err := collection.Doc(company.ID).Set(ctx, company)
	return err
}
