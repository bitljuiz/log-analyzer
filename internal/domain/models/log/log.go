package log

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Record это обертка над логом, которая содержит токенизированную информацию о каждой части лога.
type Record struct {
	Addr      string
	User      string
	Date      DateFormat
	Request   RequestFormat
	Status    HTTPStatus
	Bytes     int
	Referer   string
	UserAgent string
}

const (
	logFormat = `^\s*(?P<addr>\S+) - (?P<user>\S+) \[(?P<date>[^\]]+)\]` +
		` "(?P<request>\S+ \S+ \S+)" (?P<status>\d{3}) (?P<bytes>\d+|-) "(?P<referer>[^"]*)" "(?P<user_agent>[^"]*)"`
)

var (
	ErrInvalidLog = errors.New("invalid log")
)

func tokenizeAndParseLog(re *regexp.Regexp, matches []string) (*Record, bool, error) {
	addr := matches[re.SubexpIndex("addr")]

	if err := Validate(addr); err != nil {
		return nil, false, err
	}

	user := matches[re.SubexpIndex("user")]

	date, _ := ParseDate(matches[re.SubexpIndex("date")])

	req, _ := ParseRequest(matches[re.SubexpIndex("request")])

	statusCode, _ := strconv.Atoi(matches[re.SubexpIndex("status")])

	status, err := ParseHTTPStatus(statusCode)
	if err != nil {
		return nil, false, err
	}

	bytes, err := ParseBytes(matches[re.SubexpIndex("bytes")])
	if err != nil {
		return nil, false, err
	}

	referer := matches[re.SubexpIndex("referer")]
	userAgent := matches[re.SubexpIndex("user_agent")]

	return &Record{
		Addr:      addr,
		User:      user,
		Date:      date,
		Request:   req,
		Status:    status,
		Bytes:     bytes,
		Referer:   referer,
		UserAgent: userAgent,
	}, true, nil
}

// New используется для создания обертки над логом для дальнейшей обработки
// Принимает на вход сам лог (log), а также поле (field) и паттерн (pattern), в случае
// Если были вызваны флаги filter-field и filter-value. Также New возвращает true,
// Если log проходит фильтрацию через filter-field и filter-value, и false в обратном случае
// а также error, если что-то пошло не так.
func New(log, field, pattern string) (*Record, bool, error) {
	re := regexp.MustCompile(logFormat)

	if !re.MatchString(log) {
		return nil, false, ErrInvalidLog
	}

	matches := re.FindStringSubmatch(log)

	if len(matches) != 9 {
		return nil, false, ErrInvalidLog
	}

	if index := re.SubexpIndex(field); index != -1 {
		if !strings.Contains(matches[index], pattern) {
			return nil, false, nil
		}
	}

	return tokenizeAndParseLog(re, matches)
}
