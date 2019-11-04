package ipify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type response struct {
	IP string `json:"ip"`
}

func get(uri string) (*response, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := new(response)
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetIpv4 returns the public IPv4 address of the system (outside of NAT)
func GetIpv4() (string, error) {
	r, err := get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	return r.IP, nil
}

// GetIpv6 returns the public IPv6 address of the system (outside of NAT)
func GetIpv6() (string, error) {
	r, err := get("https://api6.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	return r.IP, nil
}
