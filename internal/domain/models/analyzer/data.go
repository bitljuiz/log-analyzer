package analyzer

import (
	"math/big"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/log"
)

// ResourcesCount представляет количество запросов к ресурсам.
// Хранит значения количества запросов для каждого ресурса и порядок их отображения.
type ResourcesCount struct {
	Values    map[string]int // Мапа ресурса и количества запросов.
	KeysOrder []string       // Порядок отображения ключей ресурсов.
}

// RequestsCount представляет количество запросов по коду ответа.
// Хранит значения количества запросов для каждого кода ответа и порядок их отображения.
type RequestsCount struct {
	Values    map[log.ResponseCode]int // Мапа кода ответа и количества запросов.
	KeysOrder []log.ResponseCode       // Порядок отображения кодов ответа.
}

// IPCount представляет количество запросов по IP-адресам.
// Хранит значения количества запросов для каждого IP и порядок их отображения.
type IPCount struct {
	Values    map[string]int
	KeysOrder []string
}

// Statistics содержит аналитические данные о логах запросов.
// Дополнительно реализованными статитисками являются максимальный и минимальный размер запроса
// А также статистика самых активных IP адресов.
type Statistics struct {
	Files                []string       // Список файлов, из которых были собраны данные.
	From                 string         // Начальная дата временного диапазона.
	To                   string         // Конечная дата временного диапазона.
	RequestsCount        RequestsCount  // Количество запросов по коду ответа.
	ResourcesCount       ResourcesCount // Количество запросов к ресурсам.
	IPCount              IPCount        // Количество запросов по IP-адресам.
	MaxSizeRequest       int            // Максимальный размер запроса.
	MinSizeRequest       int            // Минимальный размер запроса.
	ByteSizes            []int          // Список размеров запросов в байтах.
	TotalRequestsNumber  *big.Int       // Общее количество запросов.
	AverageRequestNumber *big.Int       // Среднее количество запросов.
	ByteSize             *big.Int       // Общий размер данных в байтах.
	Percentile           int            // Перцентиль по размеру запросов.
}
