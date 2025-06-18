package services

import (
	"context"
	"zine_platform/models"
)

// TODO build out this interface with required functions for uploading file metadata to a database of our choosing
type ZinePlatformFileDatabaseService interface {
	GetFileMetadata(context.Context, string) (*models.File, error)
	CreateFileMetadata(context.Context, models.File) (*models.File, error)
	UpdateFileMetadata(context.Context, string, models.File) (*models.File, error)
	DeleteFileMetadata(context.Context, string) error
}
