package domain

type CarbonFootprintRepository interface {
	Save(footprint *CarbonFootprint) error
	FindByID(id string) (*CarbonFootprint, error)
}
