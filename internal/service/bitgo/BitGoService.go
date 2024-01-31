package bitgo

import (
	"bytes"
	"data-fetcher/internal"
	"fmt"
	"log"
	"net/http"
	"time"
)

type bitGoService struct {
	repo     internal.PostgresRepo
	bitGoUrl string
}

func NewBitGoService(r internal.PostgresRepo, bitGoUrl string) internal.IFetchDataService {
	return &bitGoService{
		repo:     r,
		bitGoUrl: bitGoUrl,
	}
}

func (b *bitGoService) FetchDataAndSave(symbol string, token string) (string, error) {
	jsonBody := []byte(`{"email": "", "otp": "", "password": ""}`)
	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("%s%s", b.bitGoUrl, "/user/login")
	log.Printf("Retrieve data from: %s\n", requestURL)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return "", err
	}
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	log.Println(res)
	defer req.Body.Close()

	//err := b.repo.Save("TEST")
	//if err != nil {
	//	return "", err
	//}
	return "empty", nil
}
