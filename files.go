package runner

import (
    "os"
)


func IsFile(path string) (status bool) {
    info, err := os.Stat(path)
    if err != nil {
        return
    }
    return info.Mode().IsRegular()
}


func IsDir(path string) (status bool) {
    info, err := os.Stat(path)
    if err != nil {
        return
    }
    return info.IsDir()
}

