package scenario

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Request receives parameters to create a webscenario
type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  RequestParams `json:"params"`
	ID      int           `json:"id"`
	Auth    interface{}   `json:"auth"`
}

// RequestParams receive user and password
type RequestParams struct {
	Name   string               `json:"name"`
	HostID int                  `json:"hostid"`
	Steps  []RequestParamsSteps `json:"steps"`
}

// RequestParamsSteps array of healthcheck steps
type RequestParamsSteps struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Required    string `json:"required"`
	StatusCodes int    `json:"status_codes"`
	No          int    `json:"no"`
}

// Response struct for zabbix server resonse after create a webscenario
type Response struct {
	Jsonrpc string         `json:"jsonrpc"`
	Result  ResponseResult `json:"result"`
	ID      int            `json:"id"`
}

// ResponseResult array of ids from created webscenarios, useless now...
type ResponseResult struct {
	HTTPTestIDs []string `json:"httptestids"`
}

// CreateWebscenario is used to create a webscenario
func CreateWebscenario(appName string, zabbixEndpoint string, token string, urlHealthCheck string) error {
	webscenario := &Request{
		Jsonrpc: "2.0",
		Method:  "httptest.create",
		Params: RequestParams{
			Name:   appName,
			HostID: 10084,
			Steps: []RequestParamsSteps{
				{
					Name:        appName,
					URL:         urlHealthCheck,
					Required:    "\"status\": \"UP\"",
					StatusCodes: 200,
					No:          1,
				},
			},
		},
		ID:   1,
		Auth: token,
	}

	m, err := json.Marshal(webscenario)
	if err != nil {
		return err
	}

	res, err := http.Post(zabbixEndpoint, "application/json; charset=utf-8", bytes.NewBuffer(m))
	if err != nil {
		return err
	}

	var response Response
	dec := json.NewDecoder(res.Body)
	responseErr := dec.Decode(&response)
	if responseErr != nil {
		fmt.Printf("Problem happened on decode")
		fmt.Println(responseErr)
		return err
	}

	return err
}
