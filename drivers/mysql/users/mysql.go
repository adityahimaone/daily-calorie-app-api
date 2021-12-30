package users

import (
	"Daily-Calorie-App-API/businesses/users"
	"gorm.io/gorm"
)

type repositoryUsers struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) users.Repository {
	return &repositoryUsers{
		DB: db,
	}
}

func (repository repositoryUsers) Insert(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repository.DB.Create(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repository repositoryUsers) Update(id int, user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repository.DB.Model(&recordUser).Where("id = ?", id).Updates(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	if err := repository.DB.Where("id = ?", id).First(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repository repositoryUsers) FindByID(id int) (*users.Domain, error) {
	recordUser := Users{}
	if err := repository.DB.Where("id = ?", id).First(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repository repositoryUsers) FindByEmail(email string) (*users.Domain, error) {
	recordUser := Users{}
	if err := repository.DB.Where("email = ?", email).First(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repository repositoryUsers) GetAllUsers() (*[]users.Domain, error) {
	var recordUsers []Users
	if err := repository.DB.Find(&recordUsers).Error; err != nil {
		return &[]users.Domain{}, err
	}
	result := toDomainArray(recordUsers)
	return &result, nil
}

func (repository repositoryUsers) Delete(id int, user *users.Domain) (*users.Domain, error) {
	panic("implement me")
}
