package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Request receives zabbix parameters
type Request struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  RequestParams `json:"params"`
	ID      int           `json:"id"`
	Auth    interface{}   `json:"auth"`
}

// RequestParams receive user and password
type RequestParams struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

// Response receives zabbix parameters from zabbix api
type Response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      int         `json:"id"`
}

// MakeRequest is used to login
func MakeRequest(zabbixEndpoint string, zabbixUser string, zabbixPassword string) (token string, err error) {
	loginRequest := &Request{
		Jsonrpc: "2.0",
		Method:  "user.login",
		Params: RequestParams{
			User:     zabbixUser,
			Password: zabbixPassword,
		},
		ID:   1,
		Auth: nil,
	}

	m, err := json.Marshal(loginRequest)
	if err != nil {
		return "", err
	}

	res, err := http.Post(zabbixEndpoint, "application/json; charset=utf-8", bytes.NewBuffer(m))
	if err != nil {
		fmt.Printf("fodeu no post")
		return "", err
	}

	var response Response
	dec := json.NewDecoder(res.Body)
	error := dec.Decode(&response)
	if error != nil {
		fmt.Printf("Problem happened on decode")
		fmt.Println(error)
		return "", err
	}
	resToken := response.Result.(string)

	return resToken, nil
}
