package services

import "context"

// TODO this interface will define all of the functions the file storage service can perform
// defined as an interface because maybe some day we won't want to use google cloud storage anymore
type ZinePlatformFileStorageService interface {
	//TODO does a proper 'file' object get passed to this? What should this return?
	UploadFile(context.Context, string) error
	GetFile(context.Context, string) error
	UpdateFile(context.Context, string) error
	DeleteFile(context.Context, string) error
}
