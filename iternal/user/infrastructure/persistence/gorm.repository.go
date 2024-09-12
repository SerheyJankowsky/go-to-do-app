package persistence

import (
    "errors"
    "to-do-app/iternal/user/domain/model"

    "github.com/jackc/pgx/v5/pgconn"
    "gorm.io/gorm"
)

type GormUserRepository struct {
    db *gorm.DB
}

func NewGromUserRepository(db *gorm.DB) *GormUserRepository {
    return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user *model.User) error {
    result := r.db.Create(user)
    if result.Error != nil {
        var pgError *pgconn.PgError
        if errors.As(result.Error, &pgError) {
            if pgError.Code == "23505" {
                return errors.New("user already exist")
            }
        }
        return result.Error
    }
    return result.Error
}

func (r *GormUserRepository) FindByID(id uint) (*model.User, error) {
    var user model.User
    if err := r.db.Preload("Items").First(&user, id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *GormUserRepository) FindByEmail(email string) (*model.User, error) {
    var user model.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, errors.New("user not found")
    }
    return &user, nil
}

func (r *GormUserRepository) Update(user *model.User) error {
    return r.db.Save(user).Error
}

func (r *GormUserRepository) Delete(id uint) error {
    return r.db.Delete(&model.User{}, id).Error
}
