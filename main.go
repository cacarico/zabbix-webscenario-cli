package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/u1m0/zabbix-webscenario/login"
	"github.com/u1m0/zabbix-webscenario/scenario"
)

func checkArguments() (appName, urlHealthCheck string) {
	if len(os.Args) > 2 {
		appName = os.Args[1]
		urlHealthCheck = os.Args[2]
	} else {
		panic("Did not set appname or url")
	}
	return appName, urlHealthCheck
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	zabbixEndpoint := os.Getenv("ZABBIX_ENDPOINT")
	zabbixUser := os.Getenv("ZABBIX_USER")
	zabbixPassword := os.Getenv("ZABBI_PASSWORD")

	appName, urlHealthCheck := checkArguments()
	token, err := login.MakeRequest(zabbixEndpoint, zabbixUser, zabbixPassword)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	createError := scenario.CreateWebscenario(appName, zabbixEndpoint, token, urlHealthCheck)
	if createError != nil {
		fmt.Printf("Error creating webscenario")
		fmt.Println(createError)
		os.Exit(1)
	}

}
