package entity

import "github.com/abaykerimov/test_kmf/internal/domain/entity/dto"

type Rate struct {
	Name        string  `db:"TITLE"`
	Title       string  `db:"CODE"`
	Date        string  `db:"A_DATE"`
	Description float64 `db:"VALUE"`
}

func CreateRate(dt *dto.CreateDTO) *Rate {
	rate := &Rate{}

	rate.Name = dt.Name
	rate.Description = dt.Description
	rate.Date = dt.Date

	if dt.Title != "" {
		rate.Title = dt.Title
	}

	return rate
}
