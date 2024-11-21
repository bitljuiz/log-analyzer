package log

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrInvalidRequest = fmt.Errorf("invalid request")
)

// RequestFormat это обертка над "$request" в логе.
type RequestFormat struct {
	// Request хранит запрос в формате *http.Request.
	Request *http.Request
	// Protocol хранит значение протокола.
	Protocol string
}

// ParseRequest принимает часть лога, находящуюся в "$request" и возвращает обертку
// RequestFormat, а также error, если что-то пошло не так.
func ParseRequest(request string) (RequestFormat, error) {
	parts := strings.SplitN(request, " ", 3)

	if len(parts) != 3 {
		return RequestFormat{}, ErrInvalidRequest
	}

	method := parts[0]
	rawURL := parts[1]
	protocol := parts[2]

	req, err := http.NewRequest(method, rawURL, http.NoBody)
	if err != nil {
		return RequestFormat{}, err
	}

	return RequestFormat{
		Request:  req,
		Protocol: protocol,
	}, nil
}
