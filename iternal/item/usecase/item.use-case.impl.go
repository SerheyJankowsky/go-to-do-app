package usecase

import (
	"errors"
	"to-do-app/iternal/item/delivery/http/dto"
	"to-do-app/iternal/item/domain/model"
	"to-do-app/iternal/item/domain/repository"
)

type ItemUseCaseImpl struct {
    r repository.ItemRepository
}

func NewItemUseCase(r repository.ItemRepository) *ItemUseCaseImpl {
    return &ItemUseCaseImpl{r}
}

func (u *ItemUseCaseImpl) GetItem(id uint) (*model.Item, error) {
    item, err := u.r.FindByID(id)
    if err != nil {
        return nil, err
    }
    return item, nil
}

func (u *ItemUseCaseImpl) CreateItem(item *dto.CreteItemDto,userId uint) (*model.Item, error) {
    newItem:=model.Item{
        Title: item.Title,
        Description: item.Description,
        UserId: userId,
    }
    if err := u.r.Create(&newItem); err != nil {
        return nil, err
    }
    return &newItem, nil
}

func(u *ItemUseCaseImpl)GetItemByUser(userId uint)([]model.Item, error){
    return u.r.FindUserItem(userId)
}

func (u *ItemUseCaseImpl) UpdateItem(updateFields map[string]interface{},id uint,userId uint) (*model.Item, error) {
    item,err:=u.r.FindByID(id)
    if err!=nil{
        return nil,err
    }
    if item.UserId != userId{
        return nil,errors.New("not your todo")
    }
    if title, ok := updateFields["title"]; ok {
        item.Title = title.(string)
    }
    if description, ok := updateFields["description"]; ok {
        item.Description = description.(string)
    }
    if err := u.r.Update(item); err != nil {
        return nil, err
    }
    return item, nil
}

func (u *ItemUseCaseImpl) DeleteItem(id uint,userId uint) error {
    item,err:=u.r.FindByID(id)
    if err!=nil{
        return err
    }
    if item.UserId != userId{
        return errors.New("not your todo")
    }
    return u.r.Delete(id)
}
