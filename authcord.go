package authcord

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/denisbrodbeck/machineid"
)

type Authcord struct {
	APIKey   string
	ADMINKey string
}

type response struct {
	Message string `json:"Response"`
}

func NewAuthcordClient(apiKey string) *Authcord {
	return &Authcord{APIKey: apiKey}
}

func NewAuthcordAdmin(adminkey string) *Authcord {
	return &Authcord{ADMINKey: adminkey}
}

func GetHWID() string {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (a *Authcord) CheckHWID(hwid string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://api.authcord.xyz:8080/api/v1/check", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("hwid", hwid)
	q.Add("key", a.APIKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var checkHWIDResponse response
	err = json.Unmarshal(body, &checkHWIDResponse)
	if err != nil {
		return "", err
	}

	// remove "and" and percent sign from the message
	msg := checkHWIDResponse.Message
	msg = strings.Replace(msg, "and", "", -1)
	msg = strings.Replace(msg, "%", "", -1)

	// remove the brackets
	msg = strings.Trim(msg, "[]")

	return msg, nil
}

func (a *Authcord) AddHWID(hwid string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://api.authcord.xyz:8080/api/v1/add", nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("hwid", hwid)
	q.Add("key", a.ADMINKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	// remove "and" and percent sign from the message
	msg := response.Message
	msg = strings.Replace(msg, "and", "", -1)
	msg = strings.Replace(msg, "%", "", -1)

	// remove the brackets
	msg = strings.Trim(msg, "[]")

	return msg, nil
}

func (a *Authcord) DeleteHWID(hwid string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", "http://api.authcord.xyz:8080/api/v1/delete", nil)
	if err != nil {
		return "", err
	}
	q := req.URL.Query()
	q.Add("hwid", hwid)
	q.Add("key", a.ADMINKey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var deleteHWIDResponse response
	err = json.Unmarshal(body, &deleteHWIDResponse)
	if err != nil {
		return "", err
	}

	// remove "and" and percent sign from the message
	msg := deleteHWIDResponse.Message
	msg = strings.Replace(msg, "and", "", -1)
	msg = strings.Replace(msg, "%", "", -1)

	// remove the brackets
	msg = strings.Trim(msg, "[]")

	return msg, nil
}
