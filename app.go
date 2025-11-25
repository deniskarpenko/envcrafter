package main

import (
	"context"
	"env-crafter/pkg/db"
	"env-crafter/pkg/db/models"
	"fmt"
	"log"

	"github.com/glebarez/sqlite" // ← Pure Go драйвер
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type App struct {
	ctx       context.Context
	imageRepo *db.ImageRepository
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	dbPath := "./images.db"

	log.Printf("Opening database: %s", dbPath)

	dbGorm, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	log.Println("✓ Database connected successfully")

	a.imageRepo = db.NewImageRepository(dbGorm)
}

func (a *App) GetAllImages() []models.Image {
	images, err := a.imageRepo.GetAllImages()
	if err != nil {
		log.Printf("Error getting images: %v", err)
		return []models.Image{}
	}
	return images
}

func (a *App) GetTagByImageId(imageId int) []models.Tag {
	tags, err := a.imageRepo.GetTagsByImageId(imageId)
	if err != nil {
		log.Printf("Error getting tags: %v", err)
		return []models.Tag{}
	}
	return tags
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
