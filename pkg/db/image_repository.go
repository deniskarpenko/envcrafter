package db

import (
	"env-crafter/pkg/db/models"

	"gorm.io/gorm"
)

type ImageRepository struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{
		db: db,
	}
}
func (r *ImageRepository) GetAllImages() ([]models.Image, error) {
	var images []models.Image

	if err := r.db.Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}

func (r *ImageRepository) GetTagsByImageId(imageId int) ([]models.Tag, error) {
	var tags []models.Tag

	if err := r.db.Where("image_id = ?", imageId).Find(&tags).Error; err != nil {
		return nil, err
	}

	return tags, nil
}
