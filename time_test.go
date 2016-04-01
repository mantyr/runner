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

func TestParseDateStringOther(t *testing.T) {
    date := "29 марта 2016 03:51"
    parse := "dd mm yyyy HH:ii"
    format := "YYYY-mm-dd HH:ii:ss"

    val, err := ParseDateString(parse, format, date)
    if err != nil {
        if val != "0000-00-00" {
            t.Errorf("Error Zero return in ParseDateString, date %q, parse %q, format %q, value %q, err %q", date, parse, format, val, err)
        }
    } else {
        if val != "2016-03-29 03:51:00" {
            t.Errorf("Error ParseDateString, date %q, parse %q, format %q, value %q, err %q", date, parse, format, val, err)
        }
    }
}

