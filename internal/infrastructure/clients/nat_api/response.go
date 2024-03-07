package nat_api

const (
	defaultErrorMsg = "Наблюдаются ошибки на стороне системы"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func DefaultError() *Error {
	return &Error{
		Code:    httpStatusServerFail,
		Message: defaultErrorMsg,
	}
}

type GetExchangeRatesResponse struct {
	Rates []*RateItem `xml:"item"`
}

type RateItem struct {
	Name        string  `xml:"fullname"`
	Title       string  `xml:"title"`
	Description float64 `xml:"description"`
	Quant       int     `xml:"quant"`
	Index       string  `xml:"index"`
	Change      string  `xml:"change"`
}
