package infrastructure

import (
	"database/sql"

	"github.com/shirocola/carbon-manage-be/internal/domain"
)

type PostgresCarbonFootprintRepository struct {
	DB *sql.DB
}

func NewPostgresCarbonFootprintRepository(db *sql.DB) domain.CarbonFootprintRepository {
	return &PostgresCarbonFootprintRepository{DB: db}
}

func (r *PostgresCarbonFootprintRepository) Save(footprint *domain.CarbonFootprint) error {
	query := "INSERT INTO carbon_footprints (id, emission, sources, date) VALUES ($1, $2, $3, $4)"
	_, err := r.DB.Exec(query, footprint.ID, footprint.Emission, footprint.Sources, footprint.Date)
	return err
}

func (r *PostgresCarbonFootprintRepository) FindByID(id string) (*domain.CarbonFootprint, error) {
	var footprint domain.CarbonFootprint
	query := "SELECT id, emission, sources, date FROM carbon_footprints WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&footprint.ID, &footprint.Emission, &footprint.Sources, &footprint.Date)
	if err != nil {
		return nil, err
	}
	return &footprint, nil
}
