package flags

import "errors"

type FlagIota = int
type FlagType = int

const (
	Path = iota
	From
	To
	Format
	FilterField
	FilterValue
	Directory
	Filename
	Percentile
	FlagCount

	StringFlag
	IntegerFlag
)

var (
	FlagToName = map[FlagIota]string{
		Path:        "path",
		From:        "from",
		To:          "to",
		Format:      "format",
		FilterField: "filter-field",
		FilterValue: "filter-value",
		Directory:   "directory",
		Filename:    "filename",
		Percentile:  "percentile",
	}

	FlagToShorthandName = map[FlagIota]string{
		Path:        "p",
		From:        "f",
		To:          "t",
		Format:      "m",
		FilterField: "i",
		FilterValue: "a",
		Directory:   "d",
		Filename:    "n",
		Percentile:  "c",
	}

	FlagToUsage = map[FlagIota]string{
		Path:        "Set a path to processing file",
		From:        "Filters out logs that have a date later than the specified one",
		To:          "Filters out logs that have a date before than the specified one",
		Format:      "Sets an output data visual",
		FilterField: "Sets the field that would be used to filter logs",
		FilterValue: "Sets the value that would be used to filter logs (Use only with \"filter-field\")",
		Directory:   "Sets the directory where statistics will be saved",
		Filename:    "Sets the statistics output file",
		Percentile:  "Sets the percentile",
	}

	FlagToValueType = map[FlagIota]FlagType{
		Path:        StringFlag,
		From:        StringFlag,
		To:          StringFlag,
		Format:      StringFlag,
		FilterField: StringFlag,
		FilterValue: StringFlag,
		Directory:   StringFlag,
		Filename:    StringFlag,
		Percentile:  IntegerFlag,
	}

	FlagToDefaultValue = map[FlagIota]interface{}{
		Path:        "/*",
		From:        "1900-01-01",
		To:          "2050-01-31",
		Format:      "markdown",
		FilterField: "",
		FilterValue: "",
		Directory:   "",
		Filename:    "statistics",
		Percentile:  95,
	}

	ErrTypeNotProvided = errors.New("type not provided")
)

// CreateValue создаёт новое значение для указанного типа флага.
// Принимает тип значения (valueType) и тип флага (flagType).
// Возвращает объект Value, соответствующий значению по умолчанию для данного флага, либо ErrTypeNotProvided,
// Если тип не поддерживается.
func CreateValue(valueType, flagType int) (Value, error) {
	switch valueType {
	case StringFlag:
		return NewStringValue(FlagToDefaultValue[flagType].(string)), nil
	case IntegerFlag:
		return NewIntegerValue(FlagToDefaultValue[flagType].(int)), nil
	default:
		return nil, ErrTypeNotProvided
	}
}

// Create создаёт мапу всех флагов с их значениями по умолчанию.
// По сути, представляет собой фабрику флагов
// В целом, можно было обойтись и без нее, поскольку флагов не очень много, но
// Я посчитал релевантным показать, как можно сделать генерацию флагов в Go
// И построить легкомасштабируемое приложение.
func Create() (map[FlagIota]*Flag, error) {
	flags := make(map[FlagIota]*Flag)

	for flagIota := range FlagCount {
		defaultValue, err := CreateValue(FlagToValueType[flagIota], flagIota)

		if err != nil {
			return nil, err
		}

		flags[flagIota] = &Flag{
			Name:          FlagToName[flagIota],
			ShorthandName: FlagToShorthandName[flagIota],
			Use:           FlagToUsage[flagIota],
			Value:         defaultValue,
		}
	}

	return flags, nil
}
