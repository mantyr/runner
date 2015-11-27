package runner

import (
    "github.com/mantyr/gounidecode/unidecode"
    "strings"
)

func Translite(text string) string {
    text = unidecode.Unidecode(text)
    text = strings.Replace(text, " ", "_", -1)
    return text
}

func TextSlug(text string) string {
    text = Translite(text)
    text = strings.Replace(text, "'", "", -1)
    text = strings.Replace(text, ".", "", -1)
    text = strings.Replace(text, "__", "_", -1)
    text = strings.Trim(text, "_")
    text = strings.ToLower(text)
    return text
}