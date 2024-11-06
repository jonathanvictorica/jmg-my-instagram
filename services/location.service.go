package services

type ILocationService interface {
	SearchLocations(name string, top int) ([]LocationResponse, error)
}

type LocationResponse struct {
	LocationID   uint   `json:"location_id"`
	LocationName string `json:"location_name"`
}
