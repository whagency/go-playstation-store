package ps

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/whagency/go-playstation-store/v1/response"
	"io/ioutil"
	"net/http"
	"strings"
)

func (c *Client) apiRequest(url string, method string, params map[string]string) (*response.Response, error) {
	respData := &response.Response{
		StatusCode: http.StatusOK,
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.ApiUrl, url), nil)
	if err != nil {
		if c.log.Enable {
			c.log.Error.Printf("creating request error: %s", err.Error())
		}
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-psn-store-locale-override", c.Region)

	if c.log.Enable {
		c.log.Info.Printf("API REQUEST to: %s %s. query params: %v", strings.ToUpper(method), url, params)
	}

	transportCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout:   c.Timeout,
		Transport: transportCfg,
	}
	resp, err := client.Do(req)
	if err != nil {
		if c.log.Enable {
			c.log.Error.Printf("sending request error: %s", err.Error())
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if c.log.Enable {
			c.log.Error.Printf("response body error: %s", err.Error())
		}
		return nil, err
	}
	if err = json.Unmarshal(body, &respData); err != nil {
		if c.log.Enable {
			c.log.Error.Printf("response json error: %s", err.Error())
		}
		return nil, err
	}

	if len(respData.Errors) > 0 {
		respData.Message = respData.Errors[0].Message
		errorStatusCode := respData.Errors[0].StatusCode
		if errorStatusCode != 0 {
			respData.StatusCode = respData.Errors[0].StatusCode
		} else {
			respData.StatusCode = resp.StatusCode
		}
	}

	if 200 != respData.StatusCode {
		return nil, fmt.Errorf("error %d: %s", respData.StatusCode, respData.Message)
	}

	if c.log.Enable {
		c.log.Info.Printf("api code: %d; api message field: %s, errors: %v", respData.StatusCode, respData.Message, respData.Errors)
	}

	return respData, nil
}
