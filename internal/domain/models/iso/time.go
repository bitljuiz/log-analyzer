package iso

import (
	"errors"
	"time"
)

const (
	isoSimpleLayout = "2006-01-02"
)

var (
	ErrInvalidISOFormat = errors.New("invalid ISO format")
)

// ParseTime принимает строчку и преобразовывает её в time.Time
// Возвращает ErrInvalidISOFormat если преобразовать невозможно.
func ParseTime(timeString string) (time.Time, error) {
	isoWithTime, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		isoSimple, err1 := time.Parse(isoSimpleLayout, timeString)

		if err1 != nil {
			return time.Time{}, ErrInvalidISOFormat
		}

		return isoSimple, nil
	}

	return isoWithTime, nil
}
