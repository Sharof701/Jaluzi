package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"jaluzi/models"
	"jaluzi/pkg/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type productRepo struct {
	db  *pgxpool.Pool
	log logger.LoggerI
}

func NewProductRepo(db *pgxpool.Pool, log logger.LoggerI) *productRepo {
	return &productRepo{
		db:  db,
		log: log,
	}
}

func (u *productRepo) Create(ctx context.Context, req *models.ProductCreate) (*models.Product, error) {
	id := uuid.New().String()

	// To'g'irlangan SQL so'rovi
	query := `
		INSERT INTO "product" (
			id, 
			name, 
			code, 
			price, 
			product_image,
			created_at, 
			updated_at, 
			deleted_at
		) 
		VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL) 
		RETURNING id, name, code, price, product_image, created_at, updated_at, deleted_at
	`

	// Qaytariladigan ma'lumotlar uchun o'zgaruvchilar
	var (
		idd     sql.NullString
		name    sql.NullString
		code    sql.NullString
		price   sql.NullFloat64
		image   sql.NullString
		created sql.NullString
		updated sql.NullString
		deleted sql.NullString
	)

	// SQL so'rovini bajarish va natijalarni o'qish
	err := u.db.QueryRow(ctx, query, id, req.Name, req.Code, req.Price, req.ProductImage).Scan(
		&idd,
		&name,
		&code,
		&price,
		&image,
		&created,
		&updated,
		&deleted,
	)
	if err != nil {
		u.log.Error("Error while creating admin: " + err.Error())
		return nil, err
	}

	// Natijalarni qaytarish
	return &models.Product{
		Id:           idd.String,
		Name:         name.String,
		Code:         code.String,
		Price:        price.Float64,
		ProductImage: image.String,
		CreatedAt:    created.String,
		UpdatedAt:    updated.String,
		DeletedAt:    deleted.String,
	}, nil
}

func (u *productRepo) GetList(ctx context.Context, req *models.ProductGetListRequest) (*models.ProductGetListResponse, error) {
	var (
		resp   = &models.ProductGetListResponse{}
		query  string
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			code,
			price,
			product_image,
			TO_CHAR(created_at,'dd/mm/yyyy')
		FROM "product" 
		
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += offset + limit
	rows, err := u.db.Query(ctx, query)
	if err != nil {
		u.log.Error("error is while getting admin list" + err.Error())
		return nil, err
	}

	for rows.Next() {
		var (
			id         sql.NullString
			name       sql.NullString
			code       sql.NullString
			price      sql.NullFloat64
			image      sql.NullString
			created_at sql.NullString
		)

		err = rows.Scan(
			&resp.Total,
			&id,
			&name,
			&code,
			&price,
			&image,
			&created_at,
		)
		if err != nil {
			u.log.Error("error is while getting user list (scanning data)", logger.Error(err))
			return nil, err
		}

		resp.Product = append(resp.Product, &models.Product{
			Id:           id.String,
			Name:         name.String,
			Code:         code.String,
			Price:        price.Float64,
			ProductImage: image.String,
			CreatedAt:    created_at.String,
		})
	}
	return resp, nil
}

func (u *productRepo) GetByID(ctx context.Context, req *models.ProductPrimaryKey) (*models.Product, error) {
	var (
		query      string
		id         sql.NullString
		name       sql.NullString
		code       sql.NullString
		price      sql.NullFloat64
		image      sql.NullString
		created_at sql.NullString
	)

	query = `
		SELECT 
			id,
			name,
			code,			
			price,
			product_image,
			TO_CHAR(created_at,'dd/mm/yyyy')
		FROM "product" 
		WHERE id = $1

	`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&code,
		&price,
		&image,
		&created_at,
	)

	if err != nil && err.Error() != "no rows in result set" {
		u.log.Error("error while scanning data" + err.Error())
		return nil, err
	}

	return &models.Product{
		Id:           id.String,
		Name:         name.String,
		Code:         code.String,
		Price:        price.Float64,
		ProductImage: image.String,
		CreatedAt:    created_at.String,
	}, nil
}

func (u *productRepo) Update(ctx context.Context, req *models.ProductUpdate) (int64, error) {
	query := `
		UPDATE
			"product"
		SET
			name = $1,
			code = $2,
			price = $3,
			product_image = $4,
			updated_at = NOW()
		WHERE id = $5
	`

	// Parametrlarni to'g'ri tartibda uzatamiz
	args := []interface{}{
		req.Name,         // $1
		req.Code,         // $2
		req.Price,        // $3
		req.ProductImage, // $4
		req.Id,           // $5
	}

	// So‘rovni bajarish
	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		u.log.Error("error is while updating product data", logger.Error(err))
		return 0, err
	}

	// Necha qator yangilangani haqida ma'lumotni qaytaramiz
	return result.RowsAffected(), nil
}

func (u *productRepo) Delete(ctx context.Context, req *models.ProductPrimaryKey) error {

	_, err := u.db.Exec(ctx, `DELETE from product WHERE id = $1`, req.Id)
	if err != nil {
		u.log.Error("error is while deleting product", logger.Error(err))
		return err
	}

	return nil
}
