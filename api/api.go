package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

////func validateResponse(body []byte) {
//func ValidateResponse(resp APIV1Response) {
//	fmt.Println("Validating api response.")
//	if resp.GetError() != nil {
//		fmt.Println("Found an error")
//	}
//}

// SubmitRequestV1 sends HTTP request to V1 Acummulate API
func (c *Client) SubmitRequestV1(req *GenericRequest, encodeJSON bool) ([]byte, error) {
	var body []byte

	// convert parameters to JSON
	parms, err := json.Marshal(req)
	if err != nil {
		log.Println("Error parsing request parameters: ", err)
		return nil, err
	}

	// send HTTP request to API endpoint
	resp, err := http.Post(c.Server, "application/json", bytes.NewBuffer(parms))
	if err != nil {
		log.Println("Something went wrong when sending request: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	// read HTTP response
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error parsing response: ", err)
		return nil, err
	}

	// check if the response contains any errors
	bodyJSON := APIResponse{}
	if err = json.Unmarshal(body, &bodyJSON); err != nil {
		return nil, err
	}
	if bodyJSON.GetError() != nil {
		log.Println("An error was found while parsing the API response.")
		return nil, err
	}

	// get type and parse data
	result, err := bodyJSON.ParseData()
	if err != nil {
		return nil, err
	}

	// convert to []byte
	byteResult, err := json.Marshal(result)

	return byteResult, nil
}
