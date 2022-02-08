package histories

import (
	"Daily-Calorie-App-API/business/foods"
	"Daily-Calorie-App-API/drivers/databases/histories_detail"
	"time"
)

type Domain struct {
	ID                int
	UserID            int
	UserName          string
	FoodID            int
	Water             int
	TotalCalories     float64
	TotalFood         int
	Date              string
	Title             string
	ImgURL            string
	Calories          float64
	Fat               float64
	Protein           float64
	Carbs             float64
	ServingSize       float64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	HistoriesDetailID int
	HistoriesDetail   []histories_detail.HistoriesDetail
}

type Service interface {
	CreateHistoriesFromAPI(histories *Domain, food *foods.Domain) (*Domain, error)
	CreateHistories(histories *Domain) (*Domain, error)
	GetHistoriesByID(userID int) (*Domain, error)
	GetAllHistoryByUserID(userID int) (*[]Domain, error)
}

type Repository interface {
	Insert(histories *Domain) (*Domain, error)
	Update(id int, histories *Domain) (*Domain, error)
	GetHistoriesByID(id int) (*Domain, error)
	GetHistoriesByUserIDandDate(userID int, date string) (*Domain, error)
	GetAllHistoryByUserID(userID int) (*[]Domain, error)
}
