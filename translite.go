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
    r := strings.NewReplacer(
        "'", "",
        ".", "_",
        "§", "_",
        "±", "_",
        "!", "_",
        "@", "_",
        "#", "_",
        "$", "_",
        "%", "_",
        "^", "_",
        "&", "_",
        "*", "_",
        "(", "_",
        ")", "_",
        "=", "_",
        "+", "_",
        "`", "_",
        "~", "_",
        "{", "_",
        "}", "_",
        "(", "_",
        ")", "_",
        "\\", "_",
        "\"", "_",
        "|", "_",
        ",", "_",
        "?", "_",
        "__", "_",
    )

    text = Translite(text)
    text = r.Replace(text)
    text = strings.Trim(text, "_")
    text = strings.ToLower(text)
    return text
}