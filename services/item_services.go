package services

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/repositories"
)

type IItemService interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(CreateItemInput dto.CreateItemInput) (*models.Item, error)
	Update(updateId uint, update dto.UpdateItemInput) (*models.Item, error)
	Delete(deleteIdId uint) error
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}

func (s *ItemService) Create(input dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, input dto.UpdateItemInput) (*models.Item, error) {
	findedItem, err := s.FindById(itemId)
	if err != nil {
		return nil, err
	}
	if input.Name != nil {
		findedItem.Name = *input.Name
	}
	if input.Price != nil {
		findedItem.Price = *input.Price
	}
	if input.Description != nil {
		findedItem.Description = *input.Description
	}
	if input.SoldOut != nil {
		findedItem.SoldOut = *input.SoldOut
	}

	return s.repository.Update(*findedItem)
}

func (s *ItemService) Delete(deleteId uint) error {
	return s.repository.Delete(deleteId)
}
