package response

import "encoding/json"

type ProductDetailData struct {
	ProductRetrieve *ProductRetrieve `json:"productRetrieve"`
}

func (resp *ProductDetailData) LoadData(data interface{}) error {
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

type ProductRetrieve struct {
	ID              string                      `json:"id"`
	Name            string                      `json:"name"`
	Platforms       []string                    `json:"platforms"`
	ReleaseDate     string                      `json:"releaseDate"`
	PublisherName   string                      `json:"publisherName"`
	ScreenLanguages []string                    `json:"screenLanguages"`
	SpokenLanguages []string                    `json:"spokenLanguages"`
	SubType         string                      `json:"subType"`
	TopCategory     string                      `json:"topCategory"`
	Type            string                      `json:"type"`
	NpTitleId       string                      `json:"npTitleId"`
	Media           ProductRetrieveMedia        `json:"media"`
	Concept         ProductRetrieveConcept      `json:"concept"`
	Descriptions    ProductRetrieveDescriptions `json:"descriptions"`
}

type ProductRetrieveMedia []struct {
	Role string `json:"role"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

type ProductRetrieveConcept struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	PublisherName string `json:"publisherName"`
}

type ProductRetrieveDescriptions []struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
