package helper

const (
	ERR = "ERROR"
	MSG = "MESSAGE"
)

type ResponseJSON struct {
	MessageType string `json:"message_type"`
	Message string `json:"message"`
	Error bool `json:"error"`
	Data interface{} `json:"data"`
}

func NewResponseJSON(
	MessageType string,
	Message string,
	Error bool,
	Data interface{},
	) *ResponseJSON {
	return &ResponseJSON{MessageType, Message, Error, Data}
}