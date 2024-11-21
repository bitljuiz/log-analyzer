# Log Analyzer
Утилита для сборки информации о логах. Реализована на основе cobra, и может быть использована вне окружения
данного проекта.

```
✗ analyzer --help
Analyzer is the command that allows to collects information about logs and processes statistics about it

Usage:
  analyzer [flags]

Flags:
  -d, --directory string      Sets the directory where statistics will be saved
  -n, --filename string       Sets the statistics output file (default "statistics")
  -i, --filter-field string   Sets the field that would be used to filter logs
  -a, --filter-value string   Sets the value that would be used to filter logs (Use only with "filter-field")
  -m, --format string         Sets an output data visual (default "markdown")
  -f, --from string           Filters out logs that have a date later than the specified one (default "1900-01-01")
  -h, --help                  help for analyzer
  -p, --path string           Set a path to processing file (default "/*")
  -c, --percentile int        Sets the percentile (default 95)
  -t, --to string             Filters out logs that have a date before than the specified one (default "2050-01-31")
```

### Статистика

* Общая информация (дополнительно минимальный/максимальный размер лога)
* Статистика о частоте файлов 
* Статистика о частоте IP (дополнительная статистика)

### Флаги

**--directory**, *-d* — директория, в которую будет сохранена статистика в виде файла (по умолчанию текущая директория)

**--filename**, *-n* — имя файла, в котором будет сохранена статистика (по умолчанию, "statistics")

**--filter-field**, *-i* — поле, по которому будут профильтрованы логи 

**--filter-value**, *-a* — значение, по которому будут профильтрованы логи (работает только в паре с **--filter-field**)

**--format**, *-m* — формат файла, в котором будет сохранена статистика (по умолчанию "markdown")

**--from**, *-f* — фильтрует логи, оставляя только те, что произошли после указанной даты

**--to**, *-t* — фильтрует логи, оставляя только те, что произошли до указанной даты

**--path**, *-p* — путь до файла с логами, может быть Glob-паттерном или URL-ссылкой

**--percentile**, *-c* — меняет перцентиль в общей статистики (по умолчанию 95)

**--help**, *-h* — help-сообщение

### Использование 

Для того, чтобы можно было использовать утилиту вне проекта выполните из корня репозитория:

```
(cd gitfame/cmd/gitfame && go build .)
```

Как собрать приложение и установить его в `GOPATH/bin`?
```
go install ./cmd/analyzer/...
```

Чтобы вызывать установленный бинарь без указания полного пути, нужно добавить `GOPATH/bin` в `PATH`.
```
export PATH=$GOPATH/bin:$PATH
```



