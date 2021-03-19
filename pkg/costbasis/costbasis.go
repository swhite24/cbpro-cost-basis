package costbasis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/swhite24/cbpro-cost-basis/pkg/config"
)

type (
	// Info contains information on cost basis
	Info struct {
		Product          string
		Start            string
		End              string
		ProductPurchased string
		TotalCost        string
		AverageCost      string
	}
)

var client *coinbasepro.Client

func Calculate(cfg *config.Config) (*Info, error) {
	var err error
	var fills []coinbasepro.Fill
	var totalCost, totalPurchase float64
	initiateClient(cfg)
	if fills, err = getFillsFromDate(cfg.StartDate, cfg.Product); err != nil {
		return nil, err
	}

	for _, f := range fills {
		size, err := strconv.ParseFloat(f.Size, 64)
		if err != nil {
			return nil, err
		}

		price, err := strconv.ParseFloat(f.Price, 64)
		if err != nil {
			return nil, err
		}

		fee, err := strconv.ParseFloat(f.Fee, 64)
		if err != nil {
			return nil, err
		}

		totalPurchase += size
		totalCost += size*price + fee
	}

	return &Info{
		Product:          cfg.Product,
		Start:            cfg.StartDateStr,
		End:              cfg.EndDate.Format("2006-01-02"),
		ProductPurchased: fmt.Sprintf("%.8f", totalPurchase),
		TotalCost:        fmt.Sprintf("%.2f", totalCost),
		AverageCost:      fmt.Sprintf("%.2f", totalCost/totalPurchase),
	}, nil
}

func getFillsFromDate(start time.Time, product string) ([]coinbasepro.Fill, error) {
	fills := []coinbasepro.Fill{}
	cursor := client.ListFills(coinbasepro.ListFillsParams{
		ProductID: product,
	})

pagination:
	for cursor.HasMore {
		page := []coinbasepro.Fill{}
		if err := cursor.NextPage(&page); err != nil {
			return nil, err
		}
		for _, f := range page {
			if f.Side != "buy" {
				continue
			}
			if start.Before(f.CreatedAt.Time()) {
				fills = append(fills, f)
			} else {
				break pagination
			}
		}
	}

	return fills, nil
}

func initiateClient(cfg *config.Config) {
	client = coinbasepro.NewClient()
	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    cfg.BaseURL,
		Key:        cfg.Key,
		Passphrase: cfg.Passphrase,
		Secret:     cfg.Secret,
	})
}
