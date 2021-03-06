package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/scinna/CLIent/utils"
	"io/ioutil"
)

func Login(serverURL, username, password string) (string, error) {
	jsonBytes, _ := json.Marshal(struct{
		Username string
		Password string
	} { username, password })

	resp, err := client.Post(serverURL+"api/auth", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}

	data, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		var errResponse utils.ErrorResponse
		err := json.Unmarshal(data, &errResponse)
		if err != nil {
			return "", errors.New("can't read the error response! Something went wrong")
		}

		return "", errResponse
	}

	if resp.StatusCode == 502 {
		return "", errors.New("error #502: Is the server up")
	}

	var response struct {
		Name string
		Token string
	}

	err = json.Unmarshal(data, &response)
	if err != nil {
		return "", errors.New("can't read the error response! Something went wrong")
	}

	return response.Token, nil
}
