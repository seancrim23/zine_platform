package services

import "cloud.google.com/go/firestore"

// When implemented this service will encompass the CRUD operations of putting Zine data into a firestore database
type FirestoreZineService struct {
	Database *firestore.Client
}
