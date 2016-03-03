package emsdk

import "sync"

var (
	client *Client
	once   sync.Once
)

type Client struct {
	clientID     string
	clientSecret string
	baseURL      string
	adminToken   adminTokenResponse
}

func New(orgName, appName, clientID, clientSecret string) (*Client, error) {
	once.Do(func() {
		client = &Client{
			baseURL:      "https://a1.easemob.com/" + orgName + "/" + appName,
			clientID:     clientID,
			clientSecret: clientSecret,
		}

		adminToken, err := client.getAccessToken()
		if err != nil {
			return nil, err
		}

		client.adminToken = adminToken
	})

	return client, nil
}
