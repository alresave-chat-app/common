package service

type ChatRequest struct {
	UserName string `json:"userName"`
	Room     string `json:"room"`
	Message  string `json:"message"`
	DateTime int64  `json:"dateTime"`
}

type ConnectRequest struct {
	UserName string `json:"userName"`
	Room     string `json:"room"`
}
