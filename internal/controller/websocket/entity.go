package httpctrl

// reqPayload представляет структуру запроса WebSocket
type reqPayload struct {
	Action string `json:"action"`
	Url    string `json:"url,omitempty"`
}

// respErr представляет структуру для обработки ошибок
type respErr struct {
	Error string `json:"error"`
}