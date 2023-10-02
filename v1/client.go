package ps

import (
	"encoding/json"
	"github.com/whagency/go-playstation-store/v1/constants"
	"github.com/whagency/go-playstation-store/v1/response"
	"time"
)

const (
	apiBaseUrl string = "https://web.np.playstation.com/api/graphql/v1"

	methodGet  string = "GET"
	methodPost string = "POST"
	methodPut  string = "PUT"

	pageSize int = 10
)

type Config struct {
	Region   string
	Timeout  time.Duration
	Logging  uint8
	PageSize int
}

type Client struct {
	ApiUrl   string
	Region   string
	PageSize int
	Timeout  time.Duration
	log      logger
}

func NewClient(config *Config) *Client {
	if config == nil {
		config = &Config{}
	}
	client := &Client{
		ApiUrl:   apiBaseUrl,
		Region:   config.Region,
		Timeout:  config.Timeout,
		PageSize: config.PageSize,
		log:      newPsLogger(config.Logging),
	}

	return client
}

func (c *Client) GetRegions() map[string]string {
	return constants.GetRegionsList()
}

func (c *Client) GetCatalogData(category string, page int) (*response.CategoryGridRetrieve, error) {
	variables := map[string]interface{}{
		"filterBy":     []string{},
		"facetOptions": []string{},
		"id":           category,
		"pageArgs": map[string]int{
			"size":   c.PageSize,
			"offset": page*c.PageSize - c.PageSize,
		},
		"sortBy": map[string]interface{}{
			"isAscending": false,
			"name":        constants.ReleaseDate,
		},
	}

	jsonVariables, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	extensions := map[string]interface{}{
		"persistedQuery": map[string]interface{}{
			"version":    1,
			"sha256Hash": constants.CategoryGridRetrieve,
		},
	}

	jsonExtensions, err := json.Marshal(extensions)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"operationName": "categoryGridRetrieve",
		"variables":     string(jsonVariables),
		"extensions":    string(jsonExtensions),
	}

	respData, err := c.apiRequest("/op", methodGet, params)
	if err != nil {
		return nil, err
	}

	respDataFormatted := &response.CategoryData{}
	if err := respDataFormatted.LoadData(respData.Data); err != nil {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, err
	}

	return respDataFormatted.CategoryGridRetrieve, nil
}

func (c *Client) GetProductByID(productID string) (*response.ProductRetrieve, error) {
	variables := map[string]string{
		"productId": productID,
	}

	jsonVariables, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	extensions := map[string]interface{}{
		"persistedQuery": map[string]interface{}{
			"version":    1,
			"sha256Hash": constants.MetGetProductById,
		},
	}

	jsonExtensions, err := json.Marshal(extensions)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"operationName": "metGetProductById",
		"variables":     string(jsonVariables),
		"extensions":    string(jsonExtensions),
	}

	respData, err := c.apiRequest("/op", methodGet, params)
	if err != nil {
		return nil, err
	}

	respDataFormatted := &response.ProductDetailData{}
	if err := respDataFormatted.LoadData(respData.Data); err != nil {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, err
	}

	return respDataFormatted.ProductRetrieve, nil
}

func (c *Client) GetConceptByID(conceptID string) (*response.ConceptRetrieve, error) {
	variables := map[string]string{
		"conceptId": conceptID,
	}

	jsonVariables, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	extensions := map[string]interface{}{
		"persistedQuery": map[string]interface{}{
			"version":    1,
			"sha256Hash": constants.MetGetConceptById,
		},
	}

	jsonExtensions, err := json.Marshal(extensions)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"operationName": "metGetConceptById",
		"variables":     string(jsonVariables),
		"extensions":    string(jsonExtensions),
	}

	respData, err := c.apiRequest("/op", methodGet, params)
	if err != nil {
		return nil, err
	}

	respDataFormatted := &response.ConceptDetailData{}
	if err := respDataFormatted.LoadData(respData.Data); err != nil {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, err
	}

	return respDataFormatted.ConceptRetrieve, nil
}

func (c *Client) GetConceptPricingByID(conceptID string) (*response.ConceptPricingRetrieve, error) {
	variables := map[string]string{
		"conceptId": conceptID,
	}

	jsonVariables, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	extensions := map[string]interface{}{
		"persistedQuery": map[string]interface{}{
			"version":    1,
			"sha256Hash": constants.MetGetPricingDataByConceptId,
		},
	}

	jsonExtensions, err := json.Marshal(extensions)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"operationName": "metGetPricingDataByConceptId",
		"variables":     string(jsonVariables),
		"extensions":    string(jsonExtensions),
	}

	respData, err := c.apiRequest("/op", methodGet, params)
	if err != nil {
		return nil, err
	}

	respDataFormatted := &response.ConceptPricingData{}
	if err := respDataFormatted.LoadData(respData.Data); err != nil {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, err
	}

	return respDataFormatted.ConceptPricingRetrieve, nil
}

func (c *Client) GetAddOnsProductsByTitleId(titleID string) (*response.AddOnProductsByTitleIdRetrieve, error) {
	variables := map[string]interface{}{
		"npTitleId": titleID,
		"pageArgs": map[string]int{
			"size":   300,
			"offset": 0,
		},
	}

	jsonVariables, err := json.Marshal(variables)
	if err != nil {
		return nil, err
	}

	extensions := map[string]interface{}{
		"persistedQuery": map[string]interface{}{
			"version":    1,
			"sha256Hash": constants.MetGetAddOnsByTitleId,
		},
	}

	jsonExtensions, err := json.Marshal(extensions)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"operationName": "metGetAddOnsByTitleId",
		"variables":     string(jsonVariables),
		"extensions":    string(jsonExtensions),
	}

	respData, err := c.apiRequest("/op", methodGet, params)
	if err != nil {
		return nil, err
	}

	respDataFormatted := &response.AddOnsProductsData{}
	if err := respDataFormatted.LoadData(respData.Data); err != nil {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, err
	}

	return respDataFormatted.AddOnProductsByTitleIdRetrieve, nil
}

func (c *Client) GetFeatures() (*response.FeaturesRetrieve, error) {
	extensions := map[string]interface{}{
		"persistedQuery": map[string]interface{}{
			"version":    1,
			"sha256Hash": constants.FeaturesRetrieve,
		},
	}

	jsonExtensions, err := json.Marshal(extensions)
	if err != nil {
		return nil, err
	}

	params := map[string]string{
		"operationName": "featuresRetrieve",
		"extensions":    string(jsonExtensions),
	}

	respData, err := c.apiRequest("/op", methodGet, params)
	if err != nil {
		return nil, err
	}

	respDataFormatted := &response.FeaturesData{}
	if err := respDataFormatted.LoadData(respData.Data); err != nil {
		if c.log.Enable {
			c.log.Error.Printf("data format error: %s", ErrResponseFormat.Error())
		}
		return nil, err
	}

	return respDataFormatted.FeaturesRetrieve, nil
}
