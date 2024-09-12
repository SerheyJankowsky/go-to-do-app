package usecase

import (
	"to-do-app/iternal/item/delivery/http/dto"
	itemModel "to-do-app/iternal/item/domain/model"
)

type ItemUseCases interface {
    GetItem(id uint) (*itemModel.Item, error)
    GetItemByUser(id uint) ([]itemModel.Item, error)
    CreateItem(item *dto.CreteItemDto,userId uint) (*itemModel.Item, error)
    UpdateItem(item map[string]interface{},id uint,userId uint) (*itemModel.Item, error)
    DeleteItem(id uint,userId uint) error
}
