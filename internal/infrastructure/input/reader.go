package input

import (
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/log"
)

type LogReader interface {
	Read() (*log.Record, error)
}
