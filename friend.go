package emsdk

import (
	"encoding/json"
	"strings"
)

func (c *Client) AddFriend(owner, friend string) error {
	url := "users/" + owner + "/contacts/users/" + friend
	_, err := c.sendRequest(url, strings.NewReader(""), "POST")

	return err
}

func (c *Client) DeleteFriend(owner, friend string) error {
	url := "users/" + owner + "/contacts/users/" + friend
	_, err := c.sendRequest(url, strings.NewReader(""), "DELETE")

	return err
}

func (c *Client) AddBlacklist(owner string, friends []string) error {
	url := "users/" + owner + "/blocks/users/"
	request := struct {
		UserNames []string `json:"usernames"`
	}{
		UserNames: friends,
	}

	data, err := json.Marshal(request)
	if err != nil {
		return err
	}

	_, err = c.sendRequest(url, strings.NewReader(string(data)), "POST")

	return err
}

func (c *Client) DeleteBlacklist(owner, blocked string) error {
	url := "users/" + owner + "/blocks/users/" + blocked
	_, err := c.sendRequest(url, strings.NewReader(""), "DELETE")

	return err
}
