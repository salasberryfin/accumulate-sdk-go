package get

import "github.com/salasberryfin/accumulate-sdk-go/api"

// Options is implement for potential extra parameters that may be passed
type Options struct{}

// Get resource
type Get struct {
	APIClient    *api.Client
	extraOptions Options
}

// New creates a new instance of type Get
func New(apiClient *api.Client) (g *Get) {
	g = &Get{
		APIClient: apiClient,
	}

	return g
}

// FromObject retrieves information by URL
func (g Get) FromObject(url string) (resp api.GetResponse, err error) {
	req := api.GenericRequest{
		JSONRpc: "2.0",
		ID:      0,
		Method:  "get",
		Params: api.Params{
			URL: url,
		},
	}

	err = g.APIClient.SendRequestV1(req, &resp)

	return
}
