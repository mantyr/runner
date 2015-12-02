package runner

import (
    "github.com/mantyr/gounidecode/unidecode"
    "regexp"
    "strings"
)

func Translite(text string) string {
    text = unidecode.Unidecode(text)
    text = strings.Replace(text, " ", "_", -1)
    return text
}

func TextSlug(text string) string {
    r := regexp.MustCompile("[^a-zA-Z0-9]+")

    text = Translite(text)
    text = r.ReplaceAllString(text, "_")
    text = strings.Trim(text, "_")
    text = strings.ToLower(text)
    return text
}