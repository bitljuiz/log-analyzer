package file

import (
	"bufio"
	"os"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/log"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/impl"
)

// Reader реализация интерфейса input.LogReader (для чтения из файлов).
type Reader struct {
	reader         *bufio.Reader // reader для буфферизированного чтения.
	field, pattern string        // field и pattern нужны в случае фильтрации части лога по значению.
}

func NewLogReader(filepath, field, pattern string) (*Reader, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	return &Reader{
		reader:  bufio.NewReader(file),
		field:   field,
		pattern: pattern,
	}, nil
}

func (r *Reader) Read() (*log.Record, error) {
	return impl.ReadWithPattern(r.reader, r.field, r.pattern)
}
