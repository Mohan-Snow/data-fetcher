package coinmarketcap

import (
	"data-fetcher/internal"
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

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add(c.cmConfig.CoinMarketHeaderKey, token)
	req.Header.Add("Accept", "application/json")
	response, err := c.client.Do(req)
	defer response.Body.Close()
	if err != nil {
		return "", err
	}

	var objmap map[string]json.RawMessage
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(data, &objmap)
	if err != nil {
		return "", err
	}
	bs, _ := json.Marshal(&objmap)
	fmt.Println(string(bs))

	return "", nil
}
