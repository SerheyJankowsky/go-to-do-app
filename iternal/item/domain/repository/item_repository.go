package repository

import "to-do-app/iternal/item/domain/model"

type ItemRepository interface {
    Create(item *model.Item) error
    FindByID(id uint) (*model.Item, error)
    FindUserItem(id uint)([]model.Item,error)
    Update(item *model.Item) error
    Delete(id uint) error
}
