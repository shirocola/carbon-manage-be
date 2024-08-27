package usecase

import (
	"time"

	"github.com/google/uuid"
	"github.com/shirocola/carbon-manage-be/internal/domain"
)

type CarbonFootprintUsecase struct {
	repo domain.CarbonFootprintRepository
}

func NewCarbonFootprintUsecase(repo domain.CarbonFootprintRepository) *CarbonFootprintUsecase {
	return &CarbonFootprintUsecase{repo: repo}
}

func generatedUniqueID() string {
	return uuid.New().String()
}

func (uc *CarbonFootprintUsecase) CalculateAndSaveCarbonFootprint(source string, emissions float64) error {
	footprint := &domain.CarbonFootprint{
		ID:       generatedUniqueID(),
		Emission: emissions,
		Sources:  source,
		Date:     time.Now().Format("2006-01-02"),
	}
	return uc.repo.Save(footprint)
}
