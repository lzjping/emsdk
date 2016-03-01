package emsdk

var client *Client

type Client struct {
	clientID     string
	clientSecret string
	baseURL      string
	adminToken   adminTokenResponse
}

func New(orgName, appName, clientID, clientSecret string) (*Client, error) {
	if client == nil {
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
	}

	return client, nil
}
