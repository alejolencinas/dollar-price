package scraper

import (
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/alejolencinas/dollar-price/internal/config"
)

// Struct for the dollar price
type DollarPrice struct {
	Buy       string `json:"buy"`
	Sell      string `json:"sell"`
	FetchedAt string `json:"fetched_at"`
}

// Global cache
var (
	cachedPrice   *DollarPrice
	lastFetchedAt time.Time
	mu            sync.RWMutex
)

var cfg *config.Config = config.Load()

// Scraper function
func fetchDollarFromWeb() (*DollarPrice, error) {
	res, err := http.Get(cfg.BnaUrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("failed to fetch page")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	// Select first row of tbody
	firstRow := doc.Find("#billetes table.cotizacion tbody tr").First()

	buy := firstRow.Find("td").Eq(1).Text()  // 2nd td
	sell := firstRow.Find("td").Eq(2).Text() // 3rd td
	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	fetchedAt := time.Now().In(loc).Format(time.RFC3339)

	return &DollarPrice{
		Buy:       buy,
		Sell:      sell,
		FetchedAt: fetchedAt,
	}, nil
}

// Public function (with cache)
func GetDollarPrice() (*DollarPrice, error) {
	mu.RLock()
	if cachedPrice != nil && time.Since(lastFetchedAt) < 10*time.Minute {
		defer mu.RUnlock()
		return cachedPrice, nil
	}
	mu.RUnlock()

	// Cache expired → fetch new one
	mu.Lock()
	defer mu.Unlock()

	// Double check another goroutine didn’t refresh while waiting
	if cachedPrice != nil && time.Since(lastFetchedAt) < 10*time.Minute {
		return cachedPrice, nil
	}

	price, err := fetchDollarFromWeb()
	if err != nil {
		return nil, err
	}

	cachedPrice = price
	lastFetchedAt = time.Now()

	return price, nil
}
