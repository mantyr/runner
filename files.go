package runner

import (
    "os"
    "strings"
    "path/filepath"
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

//   dir/file.tar.gz -> gz
func GetFileExt(path string) (ext string){
    return strings.Trim(filepath.Ext(path), ".")
}

//   TODO dir/file.tar.gz -> tar.gz
/*
func GetFileExtAll(path string) (ext string){
    return
}
*/
