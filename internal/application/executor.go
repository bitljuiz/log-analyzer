package application

import (
	"errors"
	"math/big"
	"os"
	"sort"
	"time"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/iso"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/analyzer"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/flags"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/log"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/parser"
	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/visual"
	"github.com/spf13/cobra"
)

type FlagsMap = map[flags.FlagIota]*flags.Flag

const (
	MarkdownExtension = ".md"
	ADOCExtension     = ".adoc"
)

var (
	ErrUndefinedFlagValueType = errors.New("undefined flag value")
)

// ProcessFlags обрабатывает мапу флагов и возвращает
// Список файлов (files), фильтры (filterField и filterValue),
// Перцентиль (percentile) и error, если что-то пошло не так.
func ProcessFlags(flagsMap FlagsMap) (files []string, filterField, filterValue string, percentile int, err error) {
	path, _ := flagsMap[flags.Path].GetString()

	files, err = GetPaths(path)

	filterField, _ = flagsMap[flags.FilterField].GetString()

	filterValue, _ = flagsMap[flags.FilterValue].GetString()

	percentile, _ = flagsMap[flags.Percentile].GetInt()

	return files, filterField, filterValue, percentile, err
}

// ProcessFiles обрабатывает список файлов, применяет фильтры и собирает статистику.
func ProcessFiles(files []string, filterField, filterValue string,
	from, to time.Time, percentile int, stats *analyzer.Statistics) error {
	for _, file := range files {
		reader, err := chooseReader(file, filterField, filterValue)

		if err != nil {
			return err
		}

		if err := parser.Run(reader, from, to, stats); err != nil {
			return err
		}
	}

	stats.TotalRequestsNumber = big.NewInt(0)

	for _, cnt := range stats.ResourcesCount.Values {
		stats.TotalRequestsNumber.Add(stats.TotalRequestsNumber, big.NewInt(int64(cnt)))
	}

	stats.ResourcesCount.KeysOrder = SortMapByValues(stats.ResourcesCount.Values)
	stats.RequestsCount.KeysOrder = SortMapByValues(stats.RequestsCount.Values)
	stats.IPCount.KeysOrder = SortMapByValues(stats.IPCount.Values)

	if stats.TotalRequestsNumber.Int64() != 0 {
		stats.AverageRequestNumber = stats.ByteSize.Div(stats.ByteSize,
			stats.TotalRequestsNumber)
	} else {
		stats.AverageRequestNumber = big.NewInt(0)
	}

	sort.Ints(stats.ByteSizes)

	if len(stats.ByteSizes) > 0 {
		stats.MinSizeRequest = stats.ByteSizes[0]
		stats.MaxSizeRequest = stats.ByteSizes[len(stats.ByteSizes)-1]
	}

	stats.Percentile = Percentile(stats.ByteSizes, percentile)

	return nil
}

// WriteStatistics записывает статистику в указанный формат (Markdown или AsciiDoc).
func WriteStatistics(dir, filename, format string, stats *analyzer.Statistics) error {
	fullPath := dir + filename

	var err error

	switch format {
	case "markdown":
		err = os.WriteFile(fullPath+MarkdownExtension, visual.Markdown(stats), 0o600)
	case "adoc":
		err = os.WriteFile(fullPath+ADOCExtension, visual.ToADOC(stats), 0o600)
	}

	return err
}

// GetStatistics извлекает данные на основе флагов, выполняет обработку логов и сохраняет статистику.
func GetStatistics(flagsMap FlagsMap) error {
	fromString, _ := flagsMap[flags.From].GetString()
	toString, _ := flagsMap[flags.To].GetString()

	from, err := iso.ParseTime(fromString)
	if err != nil {
		return err
	}

	to, err := iso.ParseTime(toString)
	if err != nil {
		return err
	}

	files, filterField, filterValue, percentile, err := ProcessFlags(flagsMap)
	if err != nil {
		return err
	}

	stats := &analyzer.Statistics{
		Files: files,
		From:  fromString,
		To:    toString,
		RequestsCount: analyzer.RequestsCount{
			Values:    make(map[log.ResponseCode]int),
			KeysOrder: []log.ResponseCode{},
		},
		ResourcesCount: analyzer.ResourcesCount{
			Values:    make(map[string]int),
			KeysOrder: []string{},
		},
		IPCount: analyzer.IPCount{
			Values:    make(map[string]int),
			KeysOrder: []string{},
		},
		MinSizeRequest: 0,
		MaxSizeRequest: 0,
		ByteSizes:      make([]int, 0),
		ByteSize:       big.NewInt(0),
	}

	if err := ProcessFiles(files, filterField, filterValue, from, to, percentile, stats); err != nil {
		return err
	}

	dir, _ := flagsMap[flags.Directory].GetString()
	filename, _ := flagsMap[flags.Filename].GetString()
	format, _ := flagsMap[flags.Format].GetString()

	return WriteStatistics(dir, filename, format, stats)
}

// Run создает cobra-комманду analyzer (обертка над pflag), добавляет все флаги и запускает ее.
// Команда собирает информацию о логах и обрабатывает их статистику.
func Run() error {
	var analyzerCmd = cobra.Command{
		Use: "analyzer",
		Short: "Analyzer is the command that allows to collects information about logs and processes statistics " +
			"about it ",
		Long: "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 1 && args[0] == "help" {
				_ = cmd.Help()
				os.Exit(1)
			}
		},
	}

	flagsMap, err := flags.Create()
	if err != nil {
		return err
	}

	for _, flag := range flagsMap {
		switch flagValue := flag.Value.(type) {
		case *flags.StringValue:
			analyzerCmd.Flags().StringVarP(
				flagValue.Pointer(),
				flag.Name,
				flag.ShorthandName,
				flagValue.DefaultValue(),
				flag.Use,
			)
		case *flags.IntegerValue:
			analyzerCmd.Flags().IntVarP(
				flagValue.Pointer(),
				flag.Name,
				flag.ShorthandName,
				flagValue.DefaultValue(),
				flag.Use,
			)
		default:
			return ErrUndefinedFlagValueType
		}
	}

	if err := analyzerCmd.Execute(); err != nil {
		return err
	}

	if helpCmd, _ := analyzerCmd.Flags().GetBool("help"); helpCmd {
		return nil
	}

	return GetStatistics(flagsMap)
}
