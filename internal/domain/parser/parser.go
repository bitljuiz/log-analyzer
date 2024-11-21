package parser

import (
	"errors"
	"io"
	"math/big"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/analyzer"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/log"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/input"
)

// Run обрабатывает логи из LogReader в указанном временном диапазоне и собирает статистику.
// Возвращает error, если что-то пошло не так.
func Run(reader input.LogReader, from, to time.Time, bank *analyzer.Statistics) error {
	var (
		err       error
		logRecord *log.Record
	)

	for {
		logRecord, err = reader.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		if logRecord == nil {
			continue
		}

		formattedDate := logRecord.Date.ToTime()
		if from.After(formattedDate) || to.Before(formattedDate) {
			continue
		}

		bank.RequestsCount.Values[logRecord.Status.Code]++
		bank.ResourcesCount.Values[logRecord.Request.Request.URL.String()]++
		bank.IPCount.Values[logRecord.Addr]++

		bank.ByteSizes = append(bank.ByteSizes, logRecord.Bytes)
		bank.ByteSize.Add(bank.ByteSize, big.NewInt(int64(logRecord.Bytes)))
	}

	return nil
}
