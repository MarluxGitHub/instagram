package service

import (
	"github.com/MarluxGitHub/instagram/pkg/domain/rss/domain/models"
)

type ProviderService struct{
	provider []models.Provider
}

func (service *ProviderService) Add(provider models.Provider) models.Provider {
	service.provider = append(service.provider, provider)

	return provider
}

func (service *ProviderService) GetAll() []models.Provider {
	return service.provider
}

func (service *ProviderService) Get(id int) *models.Provider{
	return &service.provider[id]
}

func (service *ProviderService) Delete(id int) bool {
	service.provider = append(service.provider[:id], service.provider[id+1:]...)

	return true
}


func NewProviderService() ProviderService {
	return ProviderService{
		provider: make([]models.Provider, 0),
	}
}


