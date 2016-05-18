package emsdk

import (
	"encoding/json"
	"strings"
)

func (c *Client) CreateUser(username, password string) error {
	data := `{"username":"` + username + `","password":"` + password + `"}`
	_, err := c.sendRequest("users", strings.NewReader(data), "POST")

	return err
}

func (c *Client) DeleteUser(username string) error {
	url := "users/" + username
	_, err := c.sendRequest(url, strings.NewReader(""), "DELETE")

	return err
}

func (c *Client) ResetPassword(username, password string) error {
	url := "users/" + username + "/password"
	data := `{"newpassword":"` + password + `"}`
	_, err := c.sendRequest(url, strings.NewReader(data), "PUT")

	return err
}

func (c *Client) ResetNickname(username, nickname string) error {
	url := "users/" + username
	data := `{"nickname":"` + nickname + `"}`
	_, err := c.sendRequest(url, strings.NewReader(data), "PUT")
	return err
}

func (c *Client) IsOnline(username string) bool {
	url := "users/" + username + "/status"
	res, err := c.sendRequest(url, nil, "GET")

	var result map[string]interface{}
	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		return false
	}

	v, ok := result["data"].(map[string]interface{})
	if !ok {
		return false
	}

	if v[username].(string) != "online" {
		return false
	}

	return true
}
