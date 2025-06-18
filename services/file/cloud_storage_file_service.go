package services

import "context"

type CloudStorageFileStorageService struct {
	//TODO add client for file storage operations
}

// TODO implement functions for all file storage operations
// TODO does a proper 'file' object get passed to this? What should this return?
func (c *CloudStorageFileStorageService) UploadFile(context context.Context, id string) error {
	return nil
}

func (c *CloudStorageFileStorageService) GetFile(context context.Context, id string) error {
	return nil
}

func (c *CloudStorageFileStorageService) UpdateFile(context context.Context, id string) error {
	return nil
}

func (c *CloudStorageFileStorageService) DeleteFile(context context.Context, id string) error {
	return nil
}
