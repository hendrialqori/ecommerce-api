package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"internship-mini-project/internal/model"
	"io"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type ProvinceCityUseCase interface {
	FindAllProvince() (*[]model.ProvinceResponse, error)
	FindAllCityByProvincy(ProvinceID int) (*[]model.CityResponse, error)
	FindProvinceById(id int) (*model.ProvinceResponse, error)
	FindCityById(id int) (*model.CityResponse, error)
}

type ProvinceCityUseCaseImpl struct {
	Logger *logrus.Logger
}

func requestToAPI(url string, ctx context.Context) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

// FindAllCityByProvincy implements ProvinceCityUseCase.
func (p *ProvinceCityUseCaseImpl) FindAllCityByProvincy(ProvinceID int) (*[]model.CityResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/regencies/%d.json", ProvinceID)

	resBody, err := requestToAPI(url, ctx)
	if err != nil {
		p.Logger.WithError(err).Error("Error while request data API")
		return nil, err
	}

	var cities = &[]model.CityResponse{}

	if err := json.Unmarshal(resBody, cities); err != nil {
		p.Logger.WithError(err).Error("Failed to parse json")
		return nil, err
	}

	return cities, err
}

// FindAllProvince implements ProvinceCityUseCase.
func (p *ProvinceCityUseCaseImpl) FindAllProvince() (*[]model.ProvinceResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := "https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json"
	resBody, err := requestToAPI(url, ctx)
	if err != nil {
		p.Logger.WithError(err).Error("Error while request data API")
		return nil, err
	}

	var provinces = &[]model.ProvinceResponse{}

	if err := json.Unmarshal(resBody, provinces); err != nil {
		p.Logger.WithError(err).Error("Failed to parse json")
		return nil, err
	}

	return provinces, err
}

// FindProvinceById implements ProvinceCityUseCase.
func (p *ProvinceCityUseCaseImpl) FindProvinceById(id int) (*model.ProvinceResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/province/%d.json", id)
	resBody, err := requestToAPI(url, ctx)
	if err != nil {
		p.Logger.WithError(err).Error("Error while request data API")
		return nil, err
	}

	var province = &model.ProvinceResponse{}

	if err := json.Unmarshal(resBody, province); err != nil {
		p.Logger.WithError(err).Error("Failed to parse json")
		return nil, err
	}
	return province, err
}

// FindCityById implements ProvinceCityUseCase.
func (p *ProvinceCityUseCaseImpl) FindCityById(id int) (*model.CityResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url := fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/regency/%d.json", id)
	resBody, err := requestToAPI(url, ctx)
	if err != nil {
		p.Logger.WithError(err).Error("Error while request data API")
		return nil, err
	}

	var city = &model.CityResponse{}

	if err := json.Unmarshal(resBody, city); err != nil {
		p.Logger.WithError(err).Error("Failed to parse json")
		return nil, err
	}
	return city, err
}

func NewProvinceCityUseCase(logger *logrus.Logger) ProvinceCityUseCase {
	return &ProvinceCityUseCaseImpl{
		Logger: logger,
	}
}
