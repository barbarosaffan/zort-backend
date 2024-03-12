package repository

import (
	"github.com/barbarosaffan/zort-backend/database"
	"github.com/barbarosaffan/zort-backend/models"
)

func CreateZort(zort *models.Zort) error {
	return database.DB.Create(zort).Error
}

func FindAllZorts() ([]models.Zort, error) {
	var zorts []models.Zort
	err := database.DB.Find(&zorts).Error
	return zorts, err
}

func FindZortByID(id uint64) (models.Zort, error) {
	var zort models.Zort
	err := database.DB.First(&zort, id).Error
	return zort, err
}

func UpdateZort(zort *models.Zort) error {
	return database.DB.Save(zort).Error
}

func DeleteZort(zort *models.Zort) error {
	return database.DB.Delete(zort).Error
}
