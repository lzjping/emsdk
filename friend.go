package emsdk

import (
	"encoding/json"
	"strings"
)

func (c *Client) AddFriend(owner, friend string) (bool, error) {
	url := "users/" + owner + "/contacts/users/" + friend
	_, err := c.sendRequest(url, strings.NewReader(""), "POST")
	if err != nil {
		return false, err
	}
	return true, err
}

func (c *Client) DeleteFriend(owner, friend string) (bool, error) {
	url := "users/" + owner + "/contacts/users/" + friend
	_, err := c.sendRequest(url, strings.NewReader(""), "DELETE")
	if err != nil {
		return false, err
	}
	return true, err
}

func (c *Client) AddBlacklist(owner string, friends []string) (bool, error) {
	url := "users/" + owner + "/blocks/users/"
	request := struct {
		UserNames []string `json:"usernames"`
	}{
		UserNames: friends,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return false, err
	}

	_, err = c.sendRequest(url, strings.NewReader(string(data)), "POST")
	if err != nil {
		return false, err
	}
	return true, err
}

func (c *Client) DeleteBlacklist(owner, blocked string) (bool, error) {
	url := "users/" + owner + "/blocks/users/" + blocked
	_, err := c.sendRequest(url, strings.NewReader(""), "DELETE")
	if err != nil {
		return false, err
	}
	return true, err
}
