package response

import "encoding/json"

type FeaturesData struct {
	FeaturesRetrieve *FeaturesRetrieve `json:"tierSelectorOffersRetrieve"`
}

func (resp *FeaturesData) LoadData(data interface{}) error {
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

type FeaturesRetrieve struct {
	Offers Offers `json:"offers"`
}

type Offers []struct {
	ID                   string      `json:"skuId"`
	Title                string      `json:"title"`
	Description          string      `json:"description"`
	SubscriptionDuration string      `json:"subscriptionDuration"`
	Type                 string      `json:"type"`
	Price                OffersPrice `json:"price"`
}

type OffersPrice struct {
	BasePrice       string `json:"basePrice"`
	BasePriceValue  int    `json:"basePriceValue"`
	DiscountedPrice string `json:"discountedPrice"`
	DiscountedValue int    `json:"discountedValue"`
	CurrencyCode    string `json:"currencyCode"`
	DiscountText    string `json:"discountText"`
	IsFree          bool   `json:"isFree"`
}
