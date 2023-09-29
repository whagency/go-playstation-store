package response

import "encoding/json"

type CategoryData struct {
	CategoryGridRetrieve *CategoryGridRetrieve `json:"categoryGridRetrieve"`
}

func (resp *CategoryData) LoadData(data interface{}) error {
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

type CategoryGridRetrieve struct {
	ID       string           `json:"id"`
	Products CategoryProducts `json:"products"`
}

type CategoryProducts []struct {
	ID        string                `json:"id"`
	Name      string                `json:"name"`
	Platforms []string              `json:"platforms"`
	NpTitleId string                `json:"npTitleId"`
	Media     CategoryProductsMedia `json:"media"`
	Price     CategoryProductsPrice `json:"price"`
}

type CategoryProductsMedia []struct {
	Role string `json:"role"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

type CategoryProductsPrice struct {
	BasePrice       string `json:"basePrice"`
	DiscountedPrice string `json:"discountedPrice"`
	IsFree          bool   `json:"isFree"`
}
