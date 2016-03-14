package runner

import (
    "crypto/sha256"
    "encoding/hex"
    "strings"
    "os"
    "io"
)

// 65 chars sha 256
func GetHash(text string) string {
    text = Trim(text)

    hash := sha256.New()
    hash.Write([]byte(text))

    return hex.EncodeToString(hash.Sum(nil))
}

func GetHashFile(address string) string {
    file, err := os.Open(address)
    if err != nil {
        return ""
    }
    hash := sha256.New()
    io.Copy(hash, file)

    return hex.EncodeToString(hash.Sum(nil))
}

func Trim(text string) string {
    return strings.Trim(text, " \n\t\r")
}

// Remove duplicate rune, and remove first and last rune.
// Example:
// 	runner.RemoveDupRunes("string", " ", " ")                       // replace,          "  "      -> " "
// 	                                                                                     "       " -> " "
// 	runner.RemoveDupRunes("string\t\nstring", " \t\r\n", "")        // delete [ \t\r\n], return "string_string"
// 	runner.RemoveDupRunes("string\t\nstring", " \t\r\n", "_test_")  // replace,          return "string_test_string"
func RemoveDupRunes(s string, params ...string) (v string) {
    var sep, rep string
    if len(params) < 1 {
        return s
    } else if len(params) > 1 {
        sep = params[0]
        rep = params[1]
    } else {
        sep = params[0]
    }

    is_old, is := true, false

    for _, char := range s  {
        is = false

        for _, r := range sep {
            if char == r {
                is = true
            }
        }
        if !is {
            v += string(char)
        } else if !is_old {
            if rep != "" {
                v += rep
            } else {
                v += string(char)
            }
        }

        is_old = is
    }
    if is_old {
        if rep == "" {
            v = v[0:len(v)-1]
        } else {
            v = v[0:len(v)-len(rep)]
        }
    }
    return
}