package handler

import (
	"net/http"
	"zine_platform/services"
)

type File struct {
	DBService          services.ZinePlatformFileDatabaseService
	FileStorageService services.ZinePlatformFileStorageService
}

func (m *File) UploadFile(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *File) GetFile(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *File) UpdateFile(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *File) DeleteFile(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *File) CreateFileMetadata(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *File) GetFileMetadata(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *File) UpdateFileMetadata(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *File) DeleteFileMetadata(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}
