package visual

import (
	"fmt"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/log"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/analyzer"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/util"
)

type CommonInformation = map[string]string

const (
	CommonInformationMarkdownHeader = "| Metrics | Value |"
	ResourcesInformationHeader      = "| Resource | Count |"
	RequestCodesHeader              = "| Code | Name | Count |"
	IPCountHeader                   = "| IP | Count |"
	MarkdownHeader                  = "####"
)

func markdownHeader(s string) string {
	return MarkdownHeader + " " + s
}

func horizontalBar(columnCnt int) string {
	barSb := strings.Builder{}
	barSb.WriteString("|:" + strings.Repeat("-:|", columnCnt))
	barSb.WriteString(util.LineSeparator())

	return barSb.String()
}

// AddMarkdownCommonInformation добавляет таблицу общей информации в формате Markdown.
func AddMarkdownCommonInformation(sb *strings.Builder, stats *analyzer.Statistics) {
	_, _ = fmt.Fprintf(sb, "%s%s%s", markdownHeader("Common information"), util.LineSeparator(), util.LineSeparator())

	_, _ = fmt.Fprintf(sb, "%s%s", CommonInformationMarkdownHeader, util.LineSeparator())
	_, _ = fmt.Fprintf(sb, "%s", horizontalBar(2))

	common := OutputToCommon(stats)

	for _, metric := range commonInformationOrder {
		_, _ = fmt.Fprintf(sb, "| %s | %s |%s", metric, common[metric], util.LineSeparator())
	}

	_, _ = fmt.Fprintf(sb, "%s", util.LineSeparator())
}

// AddMarkdownResources добавляет в вывод таблицу ресурсов в формате markdown.
func AddMarkdownResources(sb *strings.Builder, stats *analyzer.Statistics) {
	_, _ = fmt.Fprintf(sb, "%s%s%s", markdownHeader("Resources"), util.LineSeparator(), util.LineSeparator())

	_, _ = fmt.Fprintf(sb, "%s%s", ResourcesInformationHeader, util.LineSeparator())
	_, _ = fmt.Fprintf(sb, "%s", horizontalBar(2))

	for _, resource := range stats.ResourcesCount.KeysOrder {
		_, _ = fmt.Fprintf(sb, "| `%s` | %s |%s", resource,
			FormatWithUnderscores(fmt.Sprintf("%d", stats.ResourcesCount.Values[resource])),
			util.LineSeparator())
	}

	_, _ = fmt.Fprintf(sb, "%s", util.LineSeparator())
}

// AddMarkdownRequestCodes добавляет таблицу кодов запросов в формате markdown.
func AddMarkdownRequestCodes(sb *strings.Builder, stats *analyzer.Statistics) {
	_, _ = fmt.Fprintf(sb, "%s%s%s", markdownHeader("Request codes"), util.LineSeparator(), util.LineSeparator())

	_, _ = fmt.Fprintf(sb, "%s%s", RequestCodesHeader, util.LineSeparator())
	_, _ = fmt.Fprintf(sb, "%s", horizontalBar(3))

	for _, code := range stats.RequestsCount.KeysOrder {
		_, _ = fmt.Fprintf(sb, "| %d | %s | %d |%s", code, log.CodeToMessage[code],
			stats.RequestsCount.Values[code],
			util.LineSeparator())
	}

	_, _ = fmt.Fprintf(sb, "%s", util.LineSeparator())
}

// AddMarkdownIPCount добавляет таблицу количества запросов по IP в формате markdown.
func AddMarkdownIPCount(sb *strings.Builder, stats *analyzer.Statistics) {
	_, _ = fmt.Fprintf(sb, "%s%s%s", markdownHeader("IP count"), util.LineSeparator(), util.LineSeparator())

	_, _ = fmt.Fprintf(sb, "%s%s", IPCountHeader, util.LineSeparator())
	_, _ = fmt.Fprintf(sb, "%s", horizontalBar(2))

	for _, ip := range stats.IPCount.KeysOrder {
		_, _ = fmt.Fprintf(sb, "| %s | %d |%s", ip, stats.IPCount.Values[ip], util.LineSeparator())
	}
}

// Markdown преобразует данные статистики в формат markdown.
func Markdown(data *analyzer.Statistics) []byte {
	markdownSb := &strings.Builder{}

	AddMarkdownCommonInformation(markdownSb, data)
	AddMarkdownResources(markdownSb, data)
	AddMarkdownRequestCodes(markdownSb, data)
	AddMarkdownIPCount(markdownSb, data)

	return []byte(markdownSb.String())
}
