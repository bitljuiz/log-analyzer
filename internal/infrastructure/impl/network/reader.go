package network

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/log"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/impl"
)

// Reader реализация интерфейса input.LogReader (для чтения по сети).
type Reader struct {
	reader         *bufio.Reader // reader для буфферизированного чтения.
	field, pattern string        // field и pattern нужны в случае фильтрации части лога по значению.
}

var (
	ErrUnexpectedCode = errors.New("unexpected code")
)

func NewReader(address, field, pattern string) (*Reader, error) {
	req, err := http.NewRequest(http.MethodGet, address, http.NoBody)

	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrUnexpectedCode
	}

	// Такое чтение необходимо только, чтобы линтер не жаловался.
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Reader{
		reader:  bufio.NewReader(bytes.NewBuffer(respBody)),
		field:   field,
		pattern: pattern,
	}, nil
}

func (r *Reader) Read() (*log.Record, error) {
	return impl.ReadWithPattern(r.reader, r.field, r.pattern)
}
