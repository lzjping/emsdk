package emsdk

import "strings"

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
