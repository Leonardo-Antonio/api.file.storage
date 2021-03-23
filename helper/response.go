package helper

const (
	ERR = "error"
	MSG = "message"
)

type Response struct {
	MessageType string `json:"message_type" xml:"message_type"`
	Message string `json:"message" xml:"message"`
	Error bool `json:"error" xml:"error"`
	Data interface{} `json:"data" xml:"data"`
}



