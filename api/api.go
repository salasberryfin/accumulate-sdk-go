package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SendRequestV1 submits a new HTTP request to the API endpoint associated with
// the Client. This function builds the request parameters and parses the JSON
// response, returning an error if anything goes wrong
func (c *Client) SendRequestV1(req GenericRequest, resp APIV1Response) (err error) {
	// convert parameters to JSON
	fmt.Println(req.Params.URL)
	params, err := json.Marshal(req)
	if err != nil {
		return
	}

	// send HTTP request to API endpoint
	httpResp, err := http.Post(c.Server, "application/json", bytes.NewBuffer(params))
	if err != nil {
		return
	}

	// decode HTTP response
	decoder := json.NewDecoder(httpResp.Body)
	err = decoder.Decode(resp)
	httpResp.Body.Close()
	if err != nil {
		return
	}

	// check if error field exists in JSON response
	err = resp.GetError()
	if err != nil {
		return
	}

	return
}
