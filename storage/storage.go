package storage

import (
	"context"
	"jaluzi/models"
)

type StorageI interface {
	Close()
	Admin() AdminI
	Product() ProductI
}

type AdminI interface {
	Create(ctx context.Context, req *models.AdminCreate) (*models.Admin, error)
	GetByID(ctx context.Context, req *models.AdminPrimaryKey) (*models.Admin, error)
	GetList(ctx context.Context, req *models.AdminGetListRequest) (*models.AdminGetListResponse, error)
	Update(ctx context.Context, req *models.AdminUpdate) (int64, error)
	Delete(ctx context.Context, req *models.AdminPrimaryKey) error
}

type ProductI interface {
	Create(ctx context.Context, req *models.ProductCreate) (*models.Product, error)
	GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error)
	GetList(ctx context.Context, req *models.ProductGetListRequest) (*models.ProductGetListResponse, error)
	Update(ctx context.Context, req *models.ProductUpdate) (int64, error)
	Delete(ctx context.Context, req *models.ProductPrimaryKey) error
}
