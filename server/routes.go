package server

import (
	"net/http"
	"zine_platform/handler"
	"zine_platform/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (z *ZinePlatformServer) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/zine", z.loadZineRoutes)
	router.Route("/file", z.loadFileRoutes)

	z.router = router
}

func (z *ZinePlatformServer) loadZineRoutes(router chi.Router) {
	zineHandler := &handler.Zine{
		Service: &services.FirestoreZineService{
			Database: z.firestoreDb,
		},
	}

	router.Post("/", zineHandler.CreateZine)
	router.Get("/{id}", zineHandler.GetZine)
	router.Put("/{id}", zineHandler.UpdateZine)
	router.Delete("/{id}", zineHandler.DeleteZine)
}

func (z *ZinePlatformServer) loadFileRoutes(router chi.Router) {
	fileHandler := &handler.File{
		DBService: &services.FirestoreFileDatabaseService{
			Database: z.firestoreDb,
		},
		FileStorageService: &services.CloudStorageFileStorageService{
			//TODO some cloud storage client that has been spun up should be here
		},
	}

	router.Post("/", fileHandler.UploadFile)
	router.Get("/{id}", fileHandler.GetFile)
	router.Put("/{id}", fileHandler.UpdateFile)
	router.Delete("/{id}", fileHandler.DeleteFile)

	router.Post("/metadata", fileHandler.CreateFileMetadata)
	router.Get("/metadata/{id}", fileHandler.GetFileMetadata)
	router.Put("/metadata/{id}", fileHandler.UpdateFileMetadata)
	router.Delete("/metadata/{id}", fileHandler.DeleteFileMetadata)
}
