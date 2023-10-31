package utils

import (
	"fmt"
	"os"
	"strconv"
)

func GetURL() string {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")

	isSsl, err := strconv.ParseBool(os.Getenv("SSL"))
	if err != nil {
		panic(fmt.Sprintf("some error"))
	}
	var https string
	if isSsl {
		https = "https://"
	} else {
		https = "http://"
	}
	var uri string
	if host == "localhost" {
		uri = https + host + ":" + port + "/"

	} else {
		uri = https + host + "/"
	}

	return uri
}
