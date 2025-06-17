package handler

import (
	"net/http"
	"zine_platform/services"
)

type Zine struct {
	Service services.ZinePlatformZineService
}

func (m *Zine) CreateZine(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *Zine) GetZine(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *Zine) UpdateZine(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}

func (m *Zine) DeleteZine(w http.ResponseWriter, r *http.Request) {
	//TODO build this out
}
