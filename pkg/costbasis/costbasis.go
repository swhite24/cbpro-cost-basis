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
		BuyCount         int
	}
)

func Calculate(client *coinbasepro.Client, cfg *config.Config) (*Info, error) {
	var err error
	var fills []coinbasepro.Fill
	var totalCost, totalPurchase float64
	if fills, err = getFillsFromDate(client, cfg.StartDate, cfg.EndDate, cfg.Product); err != nil {
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
		BuyCount:         len(fills),
	}, nil
}

func getFillsFromDate(client *coinbasepro.Client, start, end time.Time, product string) ([]coinbasepro.Fill, error) {
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
			if f.Side != "buy" || end.Before(f.CreatedAt.Time()) {
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
