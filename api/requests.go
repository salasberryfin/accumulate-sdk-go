package api

type Params struct {
	URL string `json:"url"`
}

type GenericRequest struct {
	JSONRpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}
