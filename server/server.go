package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type ZinePlatformServer struct {
	firestoreDb *firestore.Client
	//TODO add a cloud storage connection
	router http.Handler
	config Config
}

func NewZinePlatformServer(ctx context.Context, config Config) (*ZinePlatformServer, error) {
	server := &ZinePlatformServer{
		config: config,
	}

	//TODO figure out if theres a better way to do this?
	if value := os.Getenv("replace with firestore emulator host"); value != "" {
		fmt.Println("using firestore emulator: ", value)
	}

	fmt.Println(config.GCPProjectId)
	conf := &firebase.Config{ProjectID: config.GCPProjectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		fmt.Println("error making new firebase app: ", err)
		return nil, err
	}

	fmt.Println("making firestore connection...")
	database, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("error making firestore connection: ", err)
		return nil, err
	}
	server.firestoreDb = database

	server.loadRoutes()

	return server, nil
}

func (u *ZinePlatformServer) StartMailtoGeneratorServer(ctx context.Context) error {
	fmt.Println("starting zine platform server...")
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", u.config.ServerPort),
		Handler: u.router,
	}

	defer func() {
		if err := u.firestoreDb.Close(); err != nil {
			fmt.Println("failed to close firestore", err)
		}
	}()

	ch := make(chan error, 1)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
