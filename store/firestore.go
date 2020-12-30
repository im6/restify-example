package store

import (
	"os"
	"log"
	"context"
	"cloud.google.com/go/firestore"
)

func CreateSqlClient(ctx context.Context) *firestore.Client {
	projectID := os.Getenv("SQL_PROJECTID")

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// defer client.Close()
	return client
}
