package services

import (
	"context"
	"zine_platform/models"

	"cloud.google.com/go/firestore"
)

// When implemented this service will encompass the CRUD operations of putting Zine data into a firestore database
type FirestoreZineService struct {
	Database *firestore.Client
}

func (f *FirestoreZineService) GetZine(context context.Context, id string) (*models.Zine, error) {
	return nil, nil
}

func (f *FirestoreZineService) CreateZine(context context.Context, zine models.Zine) (*models.Zine, error) {
	return nil, nil
}

func (f *FirestoreZineService) UpdateZine(context context.Context, id string, zine models.Zine) (*models.Zine, error) {
	return nil, nil
}

func (f *FirestoreZineService) DeleteZine(context context.Context, id string) error {
	return nil
}
