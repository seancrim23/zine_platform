package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"zine_platform/models"
	"zine_platform/services"
	"zine_platform/utils"

	"github.com/go-chi/chi/v5"
)

type Zine struct {
	Service services.ZinePlatformZineService
}

// TODO user id should be passed for author of zine
func (z *Zine) CreateZine(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var code = 201
	var err error
	var zine models.Zine

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		code = 400
		fmt.Println(err)
		utils.RespondWithError(w, code, err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &zine)
	if err != nil {
		code = 400
		fmt.Println(err)
		utils.RespondWithError(w, code, err.Error())
		return
	}

	zn, err := z.Service.CreateZine(r.Context(), zine)
	if err != nil {
		code = 500
		fmt.Println(err)
		utils.RespondWithError(w, code, err.Error())
		return
	}

	utils.RespondWithJSON(w, code, zn)
}

func (z *Zine) GetZine(w http.ResponseWriter, r *http.Request) {
	var code = 200
	var err error

	zineId := chi.URLParam(r, "id")
	if zineId == "" {
		code = 400
		utils.RespondWithError(w, code, errors.New("no id passed to request").Error())
		return
	}

	zine, err := z.Service.GetZine(r.Context(), zineId)
	//determine what type of error and change code and return according error message
	if err != nil {
		code = 500
		utils.RespondWithError(w, code, err.Error())
		return
	}
	if zine == nil {
		code = 404
		utils.RespondWithError(w, code, errors.New("cannot find zine").Error())
		return
	}

	utils.RespondWithJSON(w, code, zine)
}

func (z *Zine) UpdateZine(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var code = 200
	var err error
	var zine models.Zine

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		code = 400
		utils.RespondWithError(w, code, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &zine)
	if err != nil {
		code = 400
		utils.RespondWithError(w, code, err.Error())
		return
	}

	zineId := chi.URLParam(r, "id")
	if zineId == "" {
		code = 400
		utils.RespondWithError(w, code, errors.New("no id passed to request").Error())
		return
	}
	//perform input validation

	zn, err := z.Service.UpdateZine(r.Context(), zineId, zine)
	if err != nil {
		code = 500
		utils.RespondWithError(w, code, err.Error())
		return
	}

	utils.RespondWithJSON(w, code, zn)
}

func (z *Zine) DeleteZine(w http.ResponseWriter, r *http.Request) {
	var code = 200
	var err error

	zineId := chi.URLParam(r, "id")
	if zineId == "" {
		code = 400
		utils.RespondWithError(w, code, errors.New("no id passed to request").Error())
		return
	}
	//perform input validation

	err = z.Service.DeleteZine(r.Context(), zineId)
	if err != nil {
		code = 500
		utils.RespondWithError(w, code, err.Error())
		return
	}

	utils.RespondWithJSON(w, code, map[string]string{"response": "deleted"})
}
