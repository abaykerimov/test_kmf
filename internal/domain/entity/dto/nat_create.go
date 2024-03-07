package dto

type CreateDTO struct {
	Name        string
	Title       string
	Date        string
	Description float64
}

func NewCreateDTO(
	name, title, date string,
	desc float64,
) *CreateDTO {
	return &CreateDTO{
		Name:        name,
		Title:       title,
		Date:        date,
		Description: desc,
	}
}
