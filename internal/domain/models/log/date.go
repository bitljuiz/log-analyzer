package log

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	monthToDayCount = map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}
	monthNumber = map[string]int{
		"Jan": 1,
		"Feb": 2,
		"Mar": 3,
		"Apr": 4,
		"May": 5,
		"Jun": 6,
		"Jul": 7,
		"Aug": 8,
		"Sep": 9,
		"Oct": 10,
		"Nov": 11,
		"Dec": 12,
	}
)

const (
	dateRegexp = `(?P<Day>[1-3][0-9]|0[1-9])\/(?P<Month>[a-zA-Z][a-z]{2})\/` +
		`(?P<Year>[1-9][0-9]{3}):(?P<Hour>[0-1][0-9]|2[0-4]):(?P<Minutes>[0-5][0-9]):(?P<Seconds>[0-5][0-9]) \+0000`
)

var (
	ErrInvalidDate        = fmt.Errorf("date is in invalid visual")
	ErrDateFormatMismatch = fmt.Errorf("date string doesn't match with required visual")
)

// DateFormat это обертка над полем "$time_local" в логе.
type DateFormat struct {
	Day     int
	Month   string
	Year    int
	Hour    int
	Minutes int
	Seconds int
}

func processDate(re *regexp.Regexp, matches []string) (DateFormat, error) {
	var (
		daysCnt int
		ok      bool
	)

	month := matches[re.SubexpIndex("Month")]

	if daysCnt, ok = monthToDayCount[month]; !ok {
		return DateFormat{}, ErrInvalidDate
	}

	year, _ := strconv.Atoi(matches[re.SubexpIndex("Year")])

	if year%4 == 0 && month == "Feb" {
		daysCnt++
	}

	day, _ := strconv.Atoi(matches[re.SubexpIndex("Day")])
	if day < 0 || day > daysCnt {
		return DateFormat{}, ErrInvalidDate
	}

	hour, _ := strconv.Atoi(matches[re.SubexpIndex("Hour")])

	minutes, _ := strconv.Atoi(matches[re.SubexpIndex("Minutes")])

	seconds, _ := strconv.Atoi(matches[re.SubexpIndex("Seconds")])

	return DateFormat{
		Day:     day,
		Month:   month,
		Year:    year,
		Hour:    hour,
		Minutes: minutes,
		Seconds: seconds,
	}, nil
}

// ParseDate преобразует строчку в обертку DateFormat
// Возвращает ErrDateFormatMismatch, если dateString не удовлетворяет формату
// Возвращает ErrInvalidDate, если сама дата не может существовать.
func ParseDate(dateString string) (DateFormat, error) {
	re := regexp.MustCompile(dateRegexp)

	if !re.MatchString(dateString) {
		return DateFormat{}, ErrDateFormatMismatch
	}

	matches := re.FindStringSubmatch(dateString)

	return processDate(re, matches)
}

func (df *DateFormat) String() string {
	return fmt.Sprintf("%02d-%s-%02d %02d:%02d:%02d",
		df.Day, df.Month, df.Year, df.Hour, df.Minutes, df.Seconds)
}

func (df *DateFormat) ToTime() time.Time {
	return time.Date(df.Year,
		time.Month(monthNumber[df.Month]), df.Day, df.Hour, df.Minutes, df.Seconds, 0, time.UTC)
}
