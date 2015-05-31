package runner

import (
    "testing"
)

func TestParseDateStringZero(t *testing.T) {
    date := "2014.05.12"
    parse := "yyyy-dd-mm"
    format := "YYYY-mm-dd"

    val, err := ParseDateString(parse, format, date)
    if err != nil {
        if val != "0000-00-00" {
            t.Errorf("Error Zero return in ParseDateString, date %q, parse %q, format %q, value %q, err %q", date, parse, format, val, err)
        }
    }
}

func TestParseDateString(t *testing.T) {
    date := "2014.05.12"
    parse := "yyyy.dd.mm"
    format := "YYYY-mm-dd"

    val, err := ParseDateString(parse, format, date)
    if err != nil {
        if val != "0000-00-00" {
            t.Errorf("Error Zero return in ParseDateString, date %q, parse %q, format %q, value %q, err %q", date, parse, format, val, err)
        }
    } else {
        if val != "2014-12-05" {
            t.Errorf("Error ParseDateString, date %q, parse %q, format %q, value %q, err %q", date, parse, format, val, err)
        }
    }
}

