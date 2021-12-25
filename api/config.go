package api

// Client holds the client configuration
type Client struct {
	Server string
}

// APIClient is used across different modules to contact the API
var APIClient Client

// NewAPIClient creates a new instance of a client to perform API calls
func NewAPIClient(serverURL string) *Client {
	c := Client{
		Server: serverURL,
	}

	return &c
}
