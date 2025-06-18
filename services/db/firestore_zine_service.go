package services

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"zine_platform/models"

	"cloud.google.com/go/firestore"
	"github.com/fatih/structs"
	"github.com/google/uuid"
)

// When implemented this service will encompass the CRUD operations of putting Zine data into a firestore database
type FirestoreZineService struct {
	Database *firestore.Client
}

func (f *FirestoreZineService) GetZine(context context.Context, id string) (*models.Zine, error) {
	dsnap, err := f.Database.Collection("Zine").Doc(id).Get(context)
	if err != nil {
		fmt.Println("error getting zine")
		fmt.Println(err)
		return nil, errors.New("error getting zine")
	}
	var z models.Zine
	err = dsnap.DataTo(&z)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error getting zine")
	}
	return &z, nil
}

func (f *FirestoreZineService) CreateZine(context context.Context, zine models.Zine) (*models.Zine, error) {
	if zine.Id == "" {
		newZineId := uuid.New()
		zine.Id = newZineId.String()
	}
	//TODO set author field equal to whoever is creating the zine
	_, err := f.Database.Collection("Zine").Doc(zine.Id).Set(context, zine)
	if err != nil {
		fmt.Println("error creating zine")
		fmt.Println(err)
		return nil, errors.New("error creating zine")
	}
	return &zine, nil
}

func (f *FirestoreZineService) UpdateZine(context context.Context, id string, zine models.Zine) (*models.Zine, error) {
	zineFirestoreUpdateData := buildZineFirestoreUpdateData(zine)
	_, err := f.Database.Collection("Zine").Doc(id).Update(context, zineFirestoreUpdateData)
	if err != nil {
		fmt.Println("error updating zine")
		fmt.Println(err)
		return nil, errors.New("error updating zine")
	}

	return &zine, nil
}

// TODO this whole function kinda sucks but works so make it better sometime
// is there a better way to get data from a request -> validate -> pass to update?
func buildZineFirestoreUpdateData(zine models.Zine) []firestore.Update {
	fireStoreUpdates := []firestore.Update{}
	//i dont like this but i dont think i have a better choice to
	//turn my go struct into something firestore will like
	zineTempGenericMap := structs.Map(zine)
	for i, v := range zineTempGenericMap {
		if strings.ToLower(i) != "id" && strings.ToLower(i) != "promoter" {
			fireStoreUpdates = append(fireStoreUpdates, firestore.Update{Path: strings.ToLower(i), Value: v})
		}
	}
	return fireStoreUpdates
}

func (f *FirestoreZineService) DeleteZine(context context.Context, id string) error {
	_, err := f.Database.Collection("Zine").Doc(id).Delete(context)
	if err != nil {
		fmt.Println("error deleting zine")
		fmt.Println(err)
		return errors.New("error deleting zine")
	}
	return nil
}
