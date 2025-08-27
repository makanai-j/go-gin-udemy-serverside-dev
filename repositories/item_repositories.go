package repositories

import (
	"errors"
	"gin-fleamarket/models"

	"gorm.io/gorm"
)

type IItemRepository interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint, userId uint) (*models.Item, error)
	Create(newItem models.Item) (*models.Item, error)
	Update(updateItem models.Item) (*models.Item, error)
	Delete(itemId uint, userId uint) error
}

type ItemMemoryRepository struct {
	items []models.Item
}

func NewItemMemoryRepository(items []models.Item) IItemRepository {
	return &ItemMemoryRepository{items: items}
}

func (r *ItemMemoryRepository) FindAll() (*[]models.Item, error) {
	return &r.items, nil
}

func (r *ItemMemoryRepository) FindById(itemId uint, userId uint) (*models.Item, error) {
	for _, item := range r.items {
		if item.ID == itemId {
			return &item, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) Create(newItem models.Item) (*models.Item, error) {
	newItem.ID = uint(len(r.items) + 1)
	r.items = append(r.items, newItem)
	return &newItem, nil
}

func (r *ItemMemoryRepository) Update(updateItem models.Item) (*models.Item, error) {
	for i, item := range r.items {
		if item.ID == updateItem.ID {
			r.items[i] = updateItem
			return &updateItem, nil
		}
	}
	return nil, errors.New("item not found")
}

func (r *ItemMemoryRepository) Delete(deleteId uint, userId uint) error {
	for i, item := range r.items {
		if item.ID == deleteId {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) IItemRepository {
	return &ItemRepository{db: db}
}

// Create implements IItemRepository.
func (i *ItemRepository) Create(newItem models.Item) (*models.Item, error) {
	result := i.db.Create(&newItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newItem, nil
}

// Delete implements IItemRepository.
func (i *ItemRepository) Delete(itemId uint, userId uint) error {
	deleteItem, err := i.FindById(itemId, userId)
	if err != nil {
		return err
	}

	result := i.db.Delete(&deleteItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindAll implements IItemRepository.
func (i *ItemRepository) FindAll() (*[]models.Item, error) {
	var items []models.Item
	result := i.db.Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return &items, nil
}

// FindById implements IItemRepository.
func (i *ItemRepository) FindById(itemId uint, userId uint) (*models.Item, error) {
	var item models.Item
	result := i.db.First(&item, "id = ? AND user_id = ?", itemId, userId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("item not found")
		}
		return nil, result.Error
	}
	return &item, nil
}

// Update implements IItemRepository.
func (i *ItemRepository) Update(updateItem models.Item) (*models.Item, error) {
	result := i.db.Save(&updateItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateItem, nil
}
