package telematic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"
)

const baseUrl = "https://telematics.transtrack.id"

func Auth(request *AuthRequest, response *AuthResponse) error {
	url := baseUrl + "/api/login"

	bodyJson, _ := json.Marshal(&request)
	bodyReader := bytes.NewReader(bodyJson)
	req, err := http.NewRequest("POST", url, bodyReader)
	req.Header.Add("Content-Type", "application/json")

	transport := Transport()
	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
	}

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("authentication failed, client not found")
	}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return err
	}
	return nil
}

func GetDevices(token string, ids []string, response *DevicesResponse) error {
	url := baseUrl + "/api/get_devices"
	param := map[string]interface{}{
		"lang":          "en",
		"user_api_hash": token,
	}

	seen := make(map[string]struct{}, len(ids))
	j := 0
	for _, v := range ids {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		ids[j] = v
		j++
	}
	ids = ids[:j]

	if len(ids) > 0 {
		for i, id := range ids {
			param[fmt.Sprintf("id[%d]", i)] = id
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	query := req.URL.Query()
	for key, value := range param {
		query.Add(key, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = query.Encode()

	transport := Transport()
	client := &http.Client{Transport: transport}
	resp, err := client.Do(req)
	if err != nil {
		return err
	} else {
		defer resp.Body.Close()
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Println("response body: ", resp.Body)
		fmt.Printf("response status: %v\n", resp.Status)
		fmt.Printf("response header: %v\n", resp.Header)
		fmt.Println("request: ", req.URL.String())
		return err
	}
	return nil
}

func Transport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
		}).DialContext,
		MaxIdleConns:        50,
		MaxIdleConnsPerHost: 50,
		TLSHandshakeTimeout: 60 * time.Second,
	}
}
