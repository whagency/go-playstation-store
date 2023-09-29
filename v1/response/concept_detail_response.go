package response

import (
	"encoding/json"
)

type ConceptDetailData struct {
	ConceptRetrieve *ConceptRetrieve `json:"conceptRetrieve"`
}

func (resp *ConceptDetailData) LoadData(data interface{}) error {
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

type ConceptRetrieve struct {
	ID             string                        `json:"id"`
	Name           string                        `json:"name"`
	PublisherName  string                        `json:"publisherName"`
	DefaultProduct ConceptRetrieveDefaultProduct `json:"defaultProduct"`
	Media          ConceptRetrieveMedia          `json:"media"`
	Descriptions   ConceptRetrieveDescriptions   `json:"descriptions"`
}

type ConceptRetrieveDefaultProduct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	NpTitleId   string `json:"npTitleId"`
	SubType     string `json:"subType"`
	TopCategory string `json:"topCategory"`
	Type        string `json:"type"`
}

type ConceptRetrieveMedia []struct {
	Role string `json:"role"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

type ConceptRetrieveDescriptions []struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
