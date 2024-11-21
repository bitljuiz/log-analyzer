package application

import (
	"net/http"
	"path/filepath"
	"sort"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/impl/file"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/impl/network"
	"github.com/es-debug/backend-academy-2024-go-template/internal/infrastructure/input"
)

type Number interface {
	~int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func chooseReader(path, field, pattern string) (input.LogReader, error) {
	if IsURL(path) {
		return network.NewReader(path, field, pattern)
	}

	return file.NewLogReader(path, field, pattern)
}

// GetPaths возвращает список путей файлов, соответствующих переданному пути.
// Если путь является URL, проверяется его доступность и он возвращается как единственный элемент списка.
// Если это glob-паттерн, возвращаются пути, соответствующие этому паттерну.
func GetPaths(path string) ([]string, error) {
	if IsURL(path) {
		if err := CheckURL(path); err != nil {
			return nil, err
		}

		return []string{path}, nil
	}

	return filepath.Glob(path)
}

// CheckURL проверяет доступность URL, отправляя HTTP-запрос методом GET.
// Возвращает ошибку, если URL недоступен или запрос не удался.
func CheckURL(url string) error {
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)

	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// IsURL проверяет, является ли переданный путь URL-адресом.
func IsURL(path string) bool {
	return strings.HasPrefix(path, "http://") ||
		strings.HasPrefix(path, "https://")
}

// SortMapByValues сортирует ключи по убыванию их значений.
func SortMapByValues[T Number, V comparable](m map[V]T) (keys []V) {
	keys = make([]V, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys
}

// Percentile вычисляет значение указанного перцентиля для переданных значений.
// Значения сортируются, и результат интерполируется между ближайшими элементами.
func Percentile(values []int, percentile int) int {
	if len(values) == 0 {
		return 0
	}

	if percentile < 0 || percentile > 100 {
		return 0
	}

	sortedData := make([]int, len(values))
	copy(sortedData, values)

	index := int(float64(percentile) / 100 * float64(len(sortedData)-1))
	lowerIndex := index
	upperIndex := lowerIndex + 1

	if upperIndex < len(sortedData) {
		weight := float64(index) - float64(lowerIndex)
		return int(float64(sortedData[lowerIndex])*(1-weight) + float64(sortedData[upperIndex])*weight)
	}

	return sortedData[lowerIndex]
}
