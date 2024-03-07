package repositories

import (
	"context"
	"fmt"
	"github.com/abaykerimov/test_kmf/internal/domain/entity"
	"github.com/abaykerimov/test_kmf/internal/infrastructure/providers/db"
	"github.com/jmoiron/sqlx"
	"log"
)

const natTableName = "R_CURRENCY"

type Repository struct {
	conn *sqlx.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		conn: db.Conn,
	}
}

func (r *Repository) GetByDate(ctx context.Context, date, code string) ([]*entity.Rate, error) {
	rates := make([]*entity.Rate, 0) // maybe create with limit

	var codeStr = "="
	if code == "" {
		codeStr = "<>"
	}

	query := fmt.Sprintf("SELECT TITLE, CODE, VALUE, A_DATE FROM %s WHERE A_DATE = $1 AND CODE %s $2", natTableName, codeStr)

	rows, err := r.conn.QueryxContext(ctx, query, date, code)
	if err != nil {
		return nil, fmt.Errorf("queryx: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		e := &entity.Rate{}

		if err = rows.StructScan(e); err != nil {
			return nil, err
		}

		rates = append(rates, e)
	}

	return rates, nil
}

func (r *Repository) Create(ctx context.Context, entity *entity.Rate) error {
	stm, err := r.conn.PrepareNamed(fmt.Sprintf("INSERT INTO %s (TITLE, CODE, VALUE, A_DATE) VALUES (:name, :title, :description, :date)", natTableName))
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err = stm.ExecContext(ctx, entity); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
