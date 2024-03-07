package nat_api

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/gobuffalo/envy"
	"log"
	"net/http"
	"time"

	"github.com/lucperkins/rek"
)

const (
	httpStatusServerFail = 500
	httpStatusBadRequest = 400
	httpStatusDelete     = 204
)

const (
	getExchangeRatePath = "/rss/get_rates.cfm"
)

var (
	ErrUnavailable = errors.New("сервис Нацбанка не доступен")
	ErrServerError = errors.New("сервис ответил ошибкой")
)

type Client struct {
	baseURI string
	timeout int
}

func NewClient() *Client {
	return &Client{
		baseURI: envy.Get("NAT_BASE_URI", ""),
		timeout: 50,
	}
}

func (c *Client) GetExchangeRates(ctx context.Context, date string) (*GetExchangeRatesResponse, error) {

	url := fmt.Sprintf("%s?fdate=%s", getExchangeRatePath, date)

	resp, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, DefaultError()
	}

	response := &GetExchangeRatesResponse{}
	if err = xml.Unmarshal(resp, &response); err != nil {
		return nil, DefaultError()
	}

	return response, nil
}

func (c *Client) do(ctx context.Context, method, url string, request interface{}) ([]byte, error) {
	var r *rek.Response
	var err error

	r, err = rek.Do(
		method,
		c.baseURI+url,
		rek.Context(ctx),
		rek.Timeout(time.Duration(c.timeout)*time.Second),
		rek.Json(request),
	)

	if err != nil {
		return nil, ErrUnavailable
	}

	if r.StatusCode() == httpStatusDelete {
		return nil, nil
	}

	bytes, err := rek.BodyAsBytes(r.Body())
	if err != nil {
		return nil, err
	}

	log.Println(fmt.Sprintf("Ответ от Нацбанка, method: %s, url: %s, status_code: %d, elapsed_time: %v", method, url, r.StatusCode(), time.Since(time.Now())))

	if err = c.checkResponse(r, bytes); err != nil {
		return nil, err
	}

	return bytes, err
}

func (c *Client) checkResponse(r *rek.Response, bytes []byte) error {
	if r.StatusCode() >= httpStatusServerFail {
		return ErrUnavailable
	}

	if r.StatusCode() >= httpStatusBadRequest && r.StatusCode() < httpStatusServerFail {
		return ErrServerError
	}

	return nil
}
