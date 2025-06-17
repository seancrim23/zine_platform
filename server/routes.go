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
