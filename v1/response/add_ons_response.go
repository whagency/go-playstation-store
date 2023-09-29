package response

import (
	"encoding/json"
)

type AddOnsProductsData struct {
	AddOnProductsByTitleIdRetrieve *AddOnProductsByTitleIdRetrieve `json:"addOnProductsByTitleIdRetrieve"`
}

func (resp *AddOnsProductsData) LoadData(data interface{}) error {
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

type AddOnProductsByTitleIdRetrieve struct {
	AddOnProducts AddOnProducts `json:"addOnProducts"`
}

type AddOnProducts []struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	SubType     string             `json:"subType"`
	TopCategory string             `json:"topCategory"`
	Type        string             `json:"type"`
	NpTitleId   string             `json:"npTitleId"`
	Price       AddOnProductsPrice `json:"price"`
}

type AddOnProductsPrice struct {
	BasePrice       string `json:"basePrice"`
	BasePriceValue  int    `json:"basePriceValue"`
	DiscountedPrice string `json:"discountedPrice"`
	DiscountedValue int    `json:"discountedValue"`
	CurrencyCode    string `json:"currencyCode"`
	DiscountText    string `json:"discountText"`
	IsFree          bool   `json:"isFree"`
}
