package impl

import (
	"jmg-my-instagram-ms/repository"
	"jmg-my-instagram-ms/services"
)

type LocationService struct {
	locationRepo repository.ILocationRepository
}

func NewLocationService(locationRepo repository.ILocationRepository) *LocationService {
	return &LocationService{locationRepo}
}

func (s *LocationService) SearchLocations(name string, top int) ([]services.LocationResponse, error) {
	locations, _ := s.locationRepo.Search(name, top)

	var result = make([]services.LocationResponse, 0)
	for _, location := range locations {
		result = append(result, services.LocationResponse{
			LocationID:   location.LocationID,
			LocationName: location.Country + " " + location.City,
		})
	}
	return result, nil
}
