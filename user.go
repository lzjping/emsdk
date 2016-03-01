package emsdk

import (
	"net/http"
	"strings"
	"time"
)

const retryInterval = 3 * time.Second

func (c *Client) CreateUser(username, password string) (bool, error) {
	data := `{"username":"` + username + `","password":"` + password + `"}`
	_, err := c.sendRequest("users", strings.NewReader(data), "POST")
	if err != nil {
		if err == ErrEM {
			info := err.(EMError)
			if info.Code == http.StatusServiceUnavailable {
				time.Sleep(retryInterval)
				_, err := c.sendRequest("users", strings.NewReader(data), "POST")
				if err != nil {
					return false, err
				}
				return true, err
			}
		}
		return false, err
	}
	return true, err
}

func (c *Client) DeleteUser(username string) (bool, error) {
	url := "users/" + username
	_, err := c.sendRequest(url, strings.NewReader(""), "DELETE")
	if err != nil {
		if err == ErrEM {
			info := err.(EMError)
			if info.Code == http.StatusServiceUnavailable {
				time.Sleep(retryInterval)
				_, err := c.sendRequest(url, strings.NewReader(""), "DELETE")
				if err != nil {
					return false, err
				}
				return true, err
			}
		}
		return false, err
	}
	return true, err
}

func (c *Client) ResetPassword(username, password string) (bool, error) {
	url := "users/" + username + "/password"
	data := `{"newpassword":"` + password + `"}`
	_, err := c.sendRequest(url, strings.NewReader(data), "PUT")
	if err != nil {
		if err == ErrEM {
			info := err.(EMError)
			if info.Code == http.StatusServiceUnavailable {
				time.Sleep(retryInterval)
				_, err := c.sendRequest(url, strings.NewReader(""), "DELETE")
				if err != nil {
					return false, err
				}
				return true, err
			}
		}
		return false, err
	}
	return true, err
}
