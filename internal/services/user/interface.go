package user

import (
	"context"
	"main/internal/entity"
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (entity.User, error)
	Create(ctx context.Context, data Create) (entity.User, error)
	Update(ctx context.Context, data Update, created_by int) (entity.User, error)
	GetAll(ctx context.Context, filer Filter, order string) ([]entity.User, int, error)
	GetById(ctx context.Context, id int) (entity.User, error)
	Delete(ctx context.Context, id int) error
	GetByEmailWithLocation(ctx context.Context, id int, lang string) (UserWithLocation, error)
	UpdateCabinet(ctx context.Context, data UpdateCabinet, id int) (entity.User, error)
}
