package service

import (
	"go-templ/infra/models"
	"go-templ/pkg/database"

	"github.com/google/uuid"
)

func CreateDocument(document models.Document) (*models.Document, error) {
	db := database.DB

	if err := db.Create(&document).Error; err != nil {
		return nil, err
	}

	return &document, nil
}

func GetDocuments(userId uuid.UUID, limit int, offset int) (*models.GetDocumentResponse, error) {
	db := database.DB

	var count int64
	var documents []models.Document
	if err := db.Where("user_id = ?", userId).Limit(limit).Offset(offset - 1).Order("due_date ASC").Find(&documents).Error; err != nil {
		return nil, err
	}

	if err := db.Model(&models.Document{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return nil, err
	}

	return &models.GetDocumentResponse{Page: offset, Limit: limit, Data: documents, Total: int(count)}, nil
}

func GetDocumentByID(id uuid.UUID) (*models.Document, error) {
	db := database.DB

	var document models.Document
	if err := db.Where("id = ?", id).First(&document).Error; err != nil {
		return nil, err
	}

	return &document, nil
}

func GetDocumentByCode(code string) (*models.Document, error) {
	db := database.DB

	var document models.Document
	if err := db.Where("code = ?", code).First(&document).Error; err != nil {
		return nil, err
	}

	return &document, nil
}

func UpdateDocument(document models.Document) (*models.Document, error) {
	db := database.DB

	if err := db.Save(&document).Error; err != nil {
		return nil, err
	}

	return &document, nil
}

func DeleteDocument(id uuid.UUID) error {
	db := database.DB

	if err := db.Where("id = ?", id).Delete(&models.Document{}).Error; err != nil {
		return err
	}

	return nil
}
