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

//   definition go run in only linux
func GetDirBin() (dir string, err error) {
    dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
    if strings.Index(dir, "/tmp/go-build") == 0 || err != nil{
        dir, err = os.Getwd()
    }
    return
}