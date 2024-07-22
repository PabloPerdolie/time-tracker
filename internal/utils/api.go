package utils

import (
	"EffectiveMobileTestTask/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Surname    string `json:"surname"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic"`
	Address    string `json:"address"`
}

func GetPersonInfo(passportSeries, passportNumber string) (person Person, err error) {
	apiUrl := config.CONFIG.APIUrl

	url := fmt.Sprintf("%s?passportSerie=%s&passportNumber=%s", apiUrl, passportSeries, passportNumber)
	resp, err := http.Get(url)
	if err != nil {
		return person, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return person, fmt.Errorf("failed to get person info, status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&person)
	if err != nil {
		return person, err
	}

	return person, nil
}
