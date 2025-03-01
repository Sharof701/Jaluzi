package postgres

import (
	"context"
	"jaluzi/config"
	"jaluzi/pkg/logger"
	"jaluzi/storage"

	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db      *pgxpool.Pool
	log     logger.LoggerI
	admin   *adminRepo
	product *productRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {
	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = 30

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}
	var loggerLevel = new(string)
	log := logger.NewLogger("app", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()
	return &store{
		db:  pgxpool,
		log: logger.NewLogger("app", *loggerLevel),
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Admin() storage.AdminI {
	if s.admin == nil {
		s.admin = &adminRepo{
			db:  s.db,
			log: s.log,
		}
	}
	return s.admin
}

func (s *store) Product() storage.ProductI {
	if s.product == nil {
		s.product = &productRepo{
			db:  s.db,
			log: s.log,
		}
	}
	return s.product
}
