package handlers

import "github.com/abaykerimov/test_kmf/internal/domain/entity"

type saveResponse struct {
	Success bool `json:"success"`
}

type rateResponse struct {
	Success bool        `json:"success"`
	Data    []*rateItem `json:"data"`
}

type rateItem struct {
	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Date        string  `json:"date"`
	Description float64 `json:"description"`
}

func GetRateResponses(rates []*entity.Rate) *rateResponse {
	response := &rateResponse{Success: true}

	for _, v := range rates {
		response.Data = append(response.Data, &rateItem{
			Name:        v.Name,
			Title:       v.Title,
			Description: v.Description,
			Date:        v.Date,
		})
	}

	return response
}
