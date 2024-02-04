package coinmarketcap

import (
	"data-fetcher/internal"
	coinmarketmapper "data-fetcher/internal/mapper/coinmarket"
	"data-fetcher/internal/model/api/coinmarket"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	coinMarketConfig "data-fetcher/config/api/coinmarketcap"
)

type coinMarketService struct {
	repo     internal.PostgresRepo
	cmConfig coinMarketConfig.CoinMarketConfig
	client   http.Client
}

func NewCoinMarketService(
	repo internal.PostgresRepo,
	cmConfig coinMarketConfig.CoinMarketConfig,
	transport *http.Transport,
) internal.IFetchDataService {
	return &coinMarketService{
		repo:     repo,
		cmConfig: cmConfig,
		client:   http.Client{Transport: transport},
	}
}

func (c *coinMarketService) FetchDataAndSave(symbols string, token string) (string, error) {
	if symbols == "" {
		return "", errors.New("symbol value must not be empty")
	}
	requestURL := fmt.Sprintf("%s%s?symbol=%s", c.cmConfig.CoinMarketUrl, "/cryptocurrency/quotes/latest", symbols)
	log.Printf("Retrieve data from: %s\n", requestURL)

	resp, err := c.doRequest(requestURL, token)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(resp.Body)
	cmDto, err := c.unmarshalResponse(&data)
	if err != nil {
		return "", err
	}

	entity, err := coinmarketmapper.MapDtoToEntity(cmDto)
	if err != nil {
		return "", err
	}

	err = c.repo.Save(entity)
	if err != nil {
		return "", err
	}
	return "Data fetched and saved", nil
}

func (c *coinMarketService) doRequest(url string, token string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(c.cmConfig.CoinMarketHeaderKey, token)
	req.Header.Add("Accept", "application/json")
	response, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *coinMarketService) unmarshalResponse(body *[]byte) (*coinmarket.CmDto, error) {
	var cmDto coinmarket.CmDto
	err := json.Unmarshal(*body, &cmDto)
	if err != nil {
		return nil, err
	}
	return &cmDto, nil
}
