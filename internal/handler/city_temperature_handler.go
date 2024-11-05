package handler

import (
	"encoding/json"
	"go-zip-code-temperature/internal/service"
	"net/http"
)

type CityTemperatureHandler struct {
	service *service.CityTemperatureService
}

func NewCityTemperatureHandler(service *service.CityTemperatureService) *CityTemperatureHandler {
	return &CityTemperatureHandler{service: service}
}

func (h CityTemperatureHandler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")
	temperature, err := h.service.GetTemperature(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temperature)
}
