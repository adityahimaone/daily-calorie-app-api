package nutritioninfo

import (
	"time"
)

type Domain struct {
	ID          int
	Calories    float64
	Carbs       float64
	Protein     float64
	Fat         float64
	ServingSize string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Repository interface {
	Create(nutritionInfo *Domain) (*Domain, error)
	Update(id int, nutritionInfo *Domain) (*Domain, error)
	Delete(id int) error
	GetByID(id int) (*Domain, error)
	GetAll() ([]Domain, error)
}
