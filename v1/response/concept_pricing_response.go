package response

import (
	"encoding/json"
)

type ConceptPricingData struct {
	ConceptPricingRetrieve *ConceptPricingRetrieve `json:"conceptRetrieve"`
}

func (resp *ConceptPricingData) LoadData(data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, resp)
	if err != nil {
		return err
	}
	return nil
}

type ConceptPricingRetrieve struct {
	ID                 string                       `json:"id"`
	DefaultProduct     ConceptPricingDefaultProduct `json:"defaultProduct"`
	SelectableProducts SelectableProducts           `json:"selectableProducts"`
}

type SelectableProducts struct {
	PurchasableProducts PurchasableProducts `json:"purchasableProducts"`
}

type ConceptPricingDefaultProduct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	NpTitleId   string `json:"npTitleId"`
	SubType     string `json:"subType"`
	TopCategory string `json:"topCategory"`
	Type        string `json:"type"`
}

type PurchasableProducts []struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	SubType     string                   `json:"subType"`
	TopCategory string                   `json:"topCategory"`
	Type        string                   `json:"type"`
	NpTitleId   string                   `json:"npTitleId"`
	Price       PurchasableProductsPrice `json:"price"`
}

type PurchasableProductsPrice struct {
	BasePrice       string `json:"basePrice"`
	BasePriceValue  int    `json:"basePriceValue"`
	DiscountedPrice string `json:"discountedPrice"`
	DiscountedValue int    `json:"discountedValue"`
	CurrencyCode    string `json:"currencyCode"`
	DiscountText    string `json:"discountText"`
	IsFree          bool   `json:"isFree"`
}
