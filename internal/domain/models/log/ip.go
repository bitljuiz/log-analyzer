package log

import (
	"fmt"
	"net"
)

var (
	ErrInvalidIP = fmt.Errorf("invalid ip")
)

// Validate принимает строчку и проверяет, можно ли ее преобразовать в IP
// Возвращает ErrInvalidIP, если ipString не может быть преобразован в IP.
func Validate(ipString string) error {
	if net.ParseIP(ipString) == nil {
		return ErrInvalidIP
	}

	return nil
}
