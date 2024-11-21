package visual

import (
	"fmt"
	"strings"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/log"

	"github.com/es-debug/backend-academy-2024-go-template/internal/domain/models/analyzer"
	"github.com/es-debug/backend-academy-2024-go-template/pkg/util"
)

const (
	CommonInformationADOCHeader    = "|Metrics |Value"
	ResourcesInformationADOCHeader = "|Resource |Count"
	RequestCodesADOCHeader         = "|Code |Name |Count"
	IPCountADOCHeader              = "|IP |Count"
	ADOCHeader                     = "===="
	ADOCTableSymbol                = "|==="
)

func adocHeader(s string) string {
	return ADOCHeader + " " + s
}

// AddADOCCommonInformation добавляет таблицу общей информации в формате AsciiDoc.
func AddADOCCommonInformation(sb *strings.Builder, stats *analyzer.Statistics) {
	_, _ = fmt.Fprintf(sb, "%s%s%s", adocHeader("Common information"), util.LineSeparator(),
		util.LineSeparator())

	_, _ = fmt.Fprintf(sb, "%s%s", ADOCTableSymbol, util.LineSeparator())
	_, _ = fmt.Fprintf(sb, "%s%s", CommonInformationADOCHeader, util.LineSeparator())

	common := OutputToCommon(stats)

	for _, metric := range commonInformationOrder {
		_, _ = fmt.Fprintf(sb, "|%s |%s%s", metric, common[metric], util.LineSeparator())
	}

	_, _ = fmt.Fprintf(sb, "%s%s", ADOCTableSymbol, util.LineSeparator())

	_, _ = fmt.Fprintf(sb, "%s", util.LineSeparator())
}

// AddADOCResources добавляет таблицу информации о ресурсах в формате AsciiDoc.
func AddADOCResources(sb *strings.Builder, stats *analyzer.Statistics) {
	_, _ = fmt.Fprintf(sb, "%s%s%s", adocHeader("Resources"), util.LineSeparator(), util.LineSeparator())

	_, _ = fmt.Fprintf(sb, "%s%s", ADOCTableSymbol, util.LineSeparator())
	_, _ = fmt.Fprintf(sb, "%s%s", ResourcesInformationADOCHeader, util.LineSeparator())

	for _, resource := range stats.ResourcesCount.KeysOrder {
		_, _ = fmt.Fprintf(sb, "|`%s` |%s%s", resource,
			FormatWithUnderscores(fmt.Sprintf("%d", stats.ResourcesCount.Values[resource])),
			util.LineSeparator())
	}

	_, _ = fmt.Fprintf(sb, "%s%s", ADOCTableSymbol, util.LineSeparator())
	_, _ = fmt.Fprintf(sb, "%s", util.LineSeparator())
}

// AddADOCRequestCodes добавляет таблицу кодов запросов в формате AsciiDoc.
func AddADOCRequestCodes(sb *strings.Builder, stats *analyzer.Statistics) {
	_, _ = fmt.Fprintf(sb, "%s%s%s", adocHeader("Request codes"), util.LineSeparator(), util.LineSeparator())

	_, _ = fmt.Fprintf(sb, "%s%s", ADOCTableSymbol, util.LineSeparator())
	_, _ = fmt.Fprintf(sb, "%s%s", RequestCodesADOCHeader, util.LineSeparator())

	for _, code := range stats.RequestsCount.KeysOrder {
		_, _ = fmt.Fprintf(sb, "|%d |%s |%d%s", code, log.CodeToMessage[code],
			stats.RequestsCount.Values[code],
			util.LineSeparator())
	}

	_, _ = fmt.Fprintf(sb, "%s%s", ADOCTableSymbol, util.LineSeparator())
}

// AddADOCIPCount добавляет таблицу количества запросов по IP в формате AsciiDoc.
func AddADOCIPCount(sb *strings.Builder, stats *analyzer.Statistics) {
	_, _ = fmt.Fprintf(sb, "%s%s%s", adocHeader("IP Count"), util.LineSeparator(), util.LineSeparator())

	_, _ = fmt.Fprintf(sb, "%s%s", ADOCTableSymbol, util.LineSeparator())
	_, _ = fmt.Fprintf(sb, "%s%s", IPCountADOCHeader, util.LineSeparator())

	for _, ip := range stats.IPCount.KeysOrder {
		_, _ = fmt.Fprintf(sb, "| %s | %d%s", ip, stats.IPCount.Values[ip], util.LineSeparator())
	}

	_, _ = fmt.Fprintf(sb, "%s%s", ADOCTableSymbol, util.LineSeparator())
}

// ToADOC преобразует статистику в формат AsciiDoc.
func ToADOC(statistics *analyzer.Statistics) []byte {
	adocSb := &strings.Builder{}

	AddADOCCommonInformation(adocSb, statistics)
	AddADOCResources(adocSb, statistics)
	AddADOCRequestCodes(adocSb, statistics)
	AddADOCIPCount(adocSb, statistics)

	return []byte(adocSb.String())
}
