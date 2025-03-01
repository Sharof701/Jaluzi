package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"jaluzi/models"
	"jaluzi/pkg/helper"
	"jaluzi/pkg/logger"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type adminRepo struct {
	db  *pgxpool.Pool
	log logger.LoggerI
}

// NewAdminRepo initializes a new instance of adminRepo
func NewAdminRepo(db *pgxpool.Pool, log logger.LoggerI) *adminRepo {
	return &adminRepo{
		db:  db,
		log: log,
	}
}

func (u *adminRepo) Create(ctx context.Context, req *models.AdminCreate) (*models.Admin, error) {
	id := uuid.New().String()

	// To'g'irlangan SQL so'rovi
	query := `
		INSERT INTO "admin" (
			id, 
			name, 
			password, 
			created_at
		) 
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP) 
		RETURNING id, name, password, created_at
	`

	// Qaytariladigan ma'lumotlar uchun o'zgaruvchilar
	var (
		idd        sql.NullString
		name       sql.NullString
		password   sql.NullString
		created_at sql.NullString
	)

	// SQL so'rovini bajarish va natijalarni o'qish
	err := u.db.QueryRow(ctx, query, id, req.Name, req.Password).Scan(
		&idd,
		&name,
		&password,
		&created_at,
	)
	if err != nil {
		u.log.Error("Error while creating admin: " + err.Error())
		return nil, err
	}

	// Natijalarni qaytarish
	return &models.Admin{
		Id:        idd.String,
		Name:      name.String,
		Password:  password.String,
		CreatedAt: created_at.String,
	}, nil
}

func (u *adminRepo) GetList(ctx context.Context, req *models.AdminGetListRequest) (*models.AdminGetListResponse, error) {
	var (
		resp   = &models.AdminGetListResponse{}
		query  string
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			password,
			TO_CHAR(created_at,'dd/mm/yyyy')
		FROM "admin" 
		
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
			password   sql.NullString
			created_at sql.NullString
		)

		err = rows.Scan(
			&resp.Total,
			&id,
			&name,
			&password,
			&created_at,
		)

		if err != nil {
			u.log.Error("error is while getting user list (scanning data)", logger.Error(err))
			return nil, err
		}

		resp.Admin = append(resp.Admin, &models.Admin{
			Id:        id.String,
			Name:      name.String,
			Password:  password.String,
			CreatedAt: created_at.String,
		})
	}
	return resp, nil
}

func (u *adminRepo) GetByID(ctx context.Context, req *models.AdminPrimaryKey) (*models.Admin, error) {
	var (
		query      string
		id         sql.NullString
		name       sql.NullString
		password   sql.NullString
		created_at sql.NullString
	)

	query = `
		SELECT 
			id,
			name,
			password,
			TO_CHAR(created_at,'dd/mm/yyyy')
		FROM "admin" 
		WHERE id = $1

	`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&password,
		&created_at,
	)

	if err != nil && err.Error() != "no rows in result set" {
		u.log.Error("error while scanning data" + err.Error())
		return nil, err
	}

	return &models.Admin{
		Id:        id.String,
		Name:      name.String,
		Password:  password.String,
		CreatedAt: created_at.String,
	}, nil
}

func (u *adminRepo) Update(ctx context.Context, req *models.AdminUpdate) (int64, error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			"admin"
		SET
			name = :name,
			password = :password,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":       req.Id,
		"name":     req.Name,
		"password": req.Password,
	}

	query, args := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		u.log.Error("error is while updating admin data", logger.Error(err))
		return 0, err
	}

	return result.RowsAffected(), nil
}
func (u *adminRepo) Delete(ctx context.Context, req *models.AdminPrimaryKey) error {

	_, err := u.db.Exec(ctx, `DELETE from admin WHERE id = $1`, req.Id)
	if err != nil {
		u.log.Error("error is while deleting admin", logger.Error(err))
		return err
	}

	return nil
}
