package lsp

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`

	// params ...
	// we will specify the params in all request types later
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  *int   `json:"id,omitempty"`

	// Result
	// Error
}

type Notification struct {
	RPC string `json:"jsonrpc"`
}
