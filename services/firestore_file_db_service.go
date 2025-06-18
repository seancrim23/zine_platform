package services

import (
	"context"
	"zine_platform/models"

	"cloud.google.com/go/firestore"
)

// when implemented this service will handle all interactions with Firestore
// TODO is there a better way to do this?
type FirestoreFileDatabaseService struct {
	Database *firestore.Client
}

// TODO implement all interface functions
func (f *FirestoreFileDatabaseService) GetFileMetadata(context context.Context, id string) (*models.File, error) {
	return nil, nil
}

func (f *FirestoreFileDatabaseService) CreateFileMetadata(context context.Context, file models.File) (*models.File, error) {
	return nil, nil
}

func (f *FirestoreFileDatabaseService) UpdateFileMetadata(context context.Context, id string, file models.File) (*models.File, error) {
	return nil, nil
}

func (f *FirestoreFileDatabaseService) DeleteFileMetadata(context context.Context, id string) error {
	return nil
}
