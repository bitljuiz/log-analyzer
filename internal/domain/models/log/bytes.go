package log

import "strconv"

// ParseBytes принимает строку и возвращает число, равняющееся числу байтов в логе.
// Возвращает error, если что-то пошло не так.
func ParseBytes(bytes string) (int, error) {
	if bytes == "-" {
		return 0, nil
	}

	return strconv.Atoi(bytes)
}
