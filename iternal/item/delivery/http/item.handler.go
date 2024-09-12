package http

import (
	"strconv"
	"to-do-app/iternal/item/delivery/http/dto"
	"to-do-app/iternal/item/usecase"

	"github.com/gofiber/fiber/v2"
)

type ItemHendler struct {
    useCase usecase.ItemUseCases
}

func NewItemHandler(u usecase.ItemUseCases) *ItemHendler {
    return &ItemHendler{u}
}

func (h *ItemHendler)CreateItem(c *fiber.Ctx)error{
    userId:=c.Locals("user_id")
    var req dto.CreteItemDto
    if err:=c.BodyParser(&req);err!=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":err.Error(),
        })
    }
    item,err:=h.useCase.CreateItem(&req,userId.(uint))
    if err!=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "item":item,
    })
}

func(h *ItemHendler) GetUserItems(c *fiber.Ctx)error{
    id := c.Locals("user_id")
    items,err:=h.useCase.GetItemByUser(id.(uint))
    if err!=nil{
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error":err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "items":items,
    })
}

func(h *ItemHendler) DeleteItem(c *fiber.Ctx)error{
    id,_:=strconv.ParseInt(c.Params("id"),10,32)
    userId:=c.Locals("user_id")
    if err:=h.useCase.DeleteItem(uint(id),userId.(uint));err!=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "message":"item was delete",
    })
}

func(h *ItemHendler) UpdateItem(c *fiber.Ctx)error{
    id,_:= strconv.ParseInt(c.Params("id"),10,32)
    userId:=c.Locals("user_id")
    var dto dto.UpdateItemDto
    if err := c.BodyParser(&dto);err!=nil{
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":err.Error(),
        })
    }
    updateFields := make(map[string]interface{})
    if dto.Description!="" {
        updateFields["description"] = dto.Description
    }
    if dto.Title!=""{
        updateFields["title"] = dto.Title
    }
    item,err:=h.useCase.UpdateItem(updateFields,uint(id),userId.(uint))
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error":err.Error(),
        })
    }
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
        "item":item,
    })
}
