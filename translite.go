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