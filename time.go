package runner

import (
    "time"
    "errors"
    "strings"
    "fmt"
)

func ParseDate(parse, date string) (time.Time, error) {
    parse = ParseDateFormat(parse)

    date = ParseDateValue(date)
    if date == "" {
        t, _ := time.Parse("2006.01.02", "0001.01.01")
        return t, errors.New("Пустая дата")
    }
    t, err := time.Parse(parse, date)
    if err != nil {
        t, _ := time.Parse("2006.01.02", "0001.01.01")
        return t, errors.New("Не удалось распарсить дату, "+parse)
    }
    return t, nil
}

// yyyy-mm-dd
func ParseDateString(parse, format, date string) (string, error) {
    date_zero := format

    parse = ParseDateFormat(parse)
    format = ParseDateFormat(format)

    date = ParseDateValue(date)
    if date == "" {
        date_zero = ParseDateFormatStringZero(date_zero)
        return date_zero, errors.New("Пустая дата")
    }
    t, err := time.Parse(parse, date)
    if err != nil {
        date_zero = ParseDateFormatStringZero(date_zero)
        return date_zero, errors.New("Не удалось распарсить дату, "+parse+", "+format)
    }
    return t.Format(format), nil
}

func DateFormat(format string, date time.Time) string {
    format = ParseDateFormat(format)
    return date.Format(format)
}

func LastDuration(format, parse string) (time.Duration, error) {
    format = ParseDateFormat(format)
    next, err := time.Parse(format, parse)
    if err != nil {
        return 0, err
    }
    this, err := time.Parse(format, time.Now().Local().Format(format))
    if err != nil {
        return 0, err
    }

    return next.Sub(this), nil
}

// see http://golang.org/src/time/format.go
func ParseDateFormat(format string) string {
    format = strings.ToLower(format)
    format = strings.Replace(format, "mm",   "01",   -1)
    format = strings.Replace(format, "dd",   "02",   -1)

    format = strings.Replace(format, "m",   "1",   -1)
    format = strings.Replace(format, "d",   "2",   -1)

    format = strings.Replace(format, "yyyy", "2006", -1)

    format = strings.Replace(format, "hh", "15", -1)
    format = strings.Replace(format, "ii", "04", -1)
    format = strings.Replace(format, "ss", "05", -1)

    format = strings.Replace(format, "h", "3", -1)
    format = strings.Replace(format, "i", "4", -1)
    format = strings.Replace(format, "s", "5", -1)

    return format
}

func ParseDateFormatStringZero(format string) string {
    format = strings.ToLower(format)
    format = strings.Replace(format, "m",   "0",   -1)
    format = strings.Replace(format, "d",   "0",   -1)
    format = strings.Replace(format, "yyyy", "0000", -1)

    format = strings.Replace(format, "h", "0", -1)
    format = strings.Replace(format, "i", "0", -1)
    format = strings.Replace(format, "s", "0", -1)
    return format
}

func ParseDateValue(date string) string {
    date = Trim(date)
    for _, months := range ShortMonths {
        for n, month := range months {
            date = strings.Replace(date, month, fmt.Sprintf("%02d", n), -1)
        }
    }
    return date
}