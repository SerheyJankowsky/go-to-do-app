package persistence

import (
	"to-do-app/iternal/item/domain/model"

	"gorm.io/gorm"
)

type GormItemRepository struct {
    db *gorm.DB
}

func NewGromUserRepository(db *gorm.DB) *GormItemRepository {
    return &GormItemRepository{db: db}
}

func (r *GormItemRepository) Create(item *model.Item) error {
    return r.db.Create(item).Error
}

func (r *GormItemRepository) FindByID(id uint) (*model.Item, error) {
    var item model.Item
    if err := r.db.First(&item, id).Error; err != nil {
        return nil, err
    }
    return &item, nil
}

func (r *GormItemRepository) FindUserItem(id uint) ([]model.Item, error) {
    var items []model.Item
    if err := r.db.Where("user_id = ?",id).Find(&items).Error; err != nil {
        return nil, err
    }
    return items,nil
}

func (r *GormItemRepository) Update(item *model.Item) error {
    return r.db.Save(item).Error
}

func (r *GormItemRepository) Delete(id uint) error {
    return r.db.Delete(&model.Item{}, id).Error
}
