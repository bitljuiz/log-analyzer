package impl

import (
	"bufio"
	"errors"
	"io"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/log"
)

func processLastLog(line, field, pattern string) (*log.Record, error) {
	lineStr, match, err2 := log.New(line, field, pattern)

	if err2 != nil {
		return nil, err2
	}

	if !match {
		return nil, nil
	}

	return lineStr, nil
}

// ReadWithPattern выделяет общую часть для реализаций input.LogReader
// Считывает строчку через reader, и возвращает *log.Record, если все успешно,
// Или error, если что-то пошло не так.
func ReadWithPattern(reader *bufio.Reader, field, pattern string) (*log.Record, error) {
	line, err := reader.ReadString('\n')

	if err != nil {
		if errors.Is(err, io.EOF) && line != "" {
			return processLastLog(line, field, pattern)
		}

		return nil, err
	}

	return processLastLog(line, field, pattern)
}
