package main

import (
	"data-fetcher/internal/repo"
	"data-fetcher/internal/service/coinmarketcap"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"

	coinMarketConfig "data-fetcher/config/api/coinmarketcap"
	postgresConfig "data-fetcher/config/postgres"
)

func main() {
	projectConfig, err := postgresConfig.NewConfig()
	if err != nil {
		log.Println("Postgres configuration error")
		log.Fatal(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		projectConfig.DbHost, projectConfig.DbPort, projectConfig.DbUsername, projectConfig.DbPassword, projectConfig.DbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("Database initializing error")
		log.Fatal(err)
	} else {
		log.Println("Established connection to database postgres")
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Println("Database ping error")
		log.Fatal(err)
	}

	cmConfig, err := coinMarketConfig.NewConfig()
	if err != nil {
		log.Println("CoinMarket configuration error")
		log.Fatal(err)
	}

	repository := repo.NewBitGoRepository(db)
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	coinMarketService := coinmarketcap.NewCoinMarketService(repository, *cmConfig, tr)
	data, err := coinMarketService.FetchDataAndSave("BTC", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
