//go:build integration

package firestore_test

import (
	"context"
	"testing"

	fs "cloud.google.com/go/firestore"
	"github.com/perebaj/credit"
	"github.com/perebaj/credit/firestore"
	"github.com/stretchr/testify/require"
)

// projectID is a mock projectID used for testing
// this const not represent a real projectID, but must be used acc
const projectID = "test-project"

func TestPingFireStore(t *testing.T) {
	c, err := fs.NewClient(context.TODO(), projectID)
	require.NoError(t, err)

	defer teardown(t, c)

	_, err = c.Collection("test").Doc("test").Set(context.TODO(), map[string]interface{}{"test": "test"})
	require.NoError(t, err)
}

func TestStorageSaveCompany(t *testing.T) {
	c, err := fs.NewClient(context.TODO(), projectID)
	require.NoError(t, err)
	defer teardown(t, c)

	storage := firestore.NewStorage(c, projectID)

	company := credit.Company{
		ID:   "test",
		Name: "test",
	}

	err = storage.SaveCompany(context.TODO(), company)
	require.NoError(t, err)

	doc, err := c.Collection("companies").Doc(company.ID).Get(context.TODO())
	require.NoError(t, err)

	var got credit.Company
	doc.DataTo(&got)
	require.Equal(t, company, got)
}

// teardown deletes all collections and documents in the firestore database
// it must be called in all tests that uses the firestore database
func teardown(t *testing.T, c *fs.Client) {
	ctx := context.TODO()
	iter := c.Collections(ctx)
	for {
		collection, err := iter.Next()
		if err != nil {
			break
		}
		iter := collection.Documents(ctx)
		for {
			doc, err := iter.Next()
			if err != nil {
				break
			}
			_, err = doc.Ref.Delete(ctx)
			require.NoError(t, err)
		}
	}
}
