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