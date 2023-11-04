package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/config"
	"net/http"
)

func HttpRequest(method string, endPoint string, body any) (response []byte, err error) {

	// Create a GET request to the Shopify API
	client := &http.Client{}
	var req *http.Request
	var reqBody []byte
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			fmt.Println("Error creating the request:", err)
			return response, err
		}
	}

	req, err = http.NewRequest(method, fmt.Sprintf("%s%s", config.Config("BASE_URL"), endPoint), bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error creating the request:", err)
		return response, err
	}
	// Set the Shopify access token in the request headers
	req.Header.Set("X-Shopify-Access-Token", config.Config("ADMIN_TOKEN"))

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending the request:", err)
		return response, err
	}
	defer resp.Body.Close()

	// Read and send the response
	response, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response:", err)
		return response, err
	}

	return response, err

}
