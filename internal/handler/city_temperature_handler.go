package handler

import (
	"encoding/json"
	"fmt"
	"go-zip-code-temperature/config"
	"go-zip-code-temperature/internal/model"
	"io"
	"net/http"
)

type CityTemperatureHandler struct {
	config *config.Config
}

func NewCityTemperatureHandler(config *config.Config) *CityTemperatureHandler {
	return &CityTemperatureHandler{config: config}
}

func (h CityTemperatureHandler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")
	cepResponse, err := http.Get(h.config.CEPServiceURL + "/" + cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer cepResponse.Body.Close()

	var address model.AddressResponse
	err = json.NewDecoder(cepResponse.Body).Decode(&address)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reqUlr := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", h.config.WeatherAPIURL, h.config.WeatherAPIKey, address.City)
	fmt.Println("reqUlr", reqUlr)
	weatherResponse, err := http.Get(reqUlr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer weatherResponse.Body.Close()
	fmt.Println("weatherResponse", weatherResponse)

	var weather model.WeatherResponse
	if weatherResponse.StatusCode != http.StatusOK {
		fmt.Println("error: received non-200 response code", weatherResponse.StatusCode)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(weatherResponse.Body)
	if err != nil {
		fmt.Println("error reading response body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println("response body:", string(body))

	err = json.Unmarshal(body, &weather)
	if err != nil {
		fmt.Println("error unmarshalling response body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	temperatureResponse := model.TemperatureResponse{
		TempC: weather.Current.TempC,
		TempF: weather.Current.TempC*1.8 + 32,
		TempK: weather.Current.TempC + 273.15,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temperatureResponse)
}
