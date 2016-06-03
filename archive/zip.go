package archive

import (
    "archive/zip"
    "os"
    "errors"
    "io"
    "io/ioutil"
)

type Zip struct {
    File       *zip.ReadCloser
    name       string
    File_list  map[string]int
    File_unzip map[string]*os.File
}

func ZipOpen(name string) (file *Zip, err error) {
    file = new(Zip)
    file.File, err = zip.OpenReader(name)
    file.name = name
    if err != nil {
        return
    }
    file.File_list = make(map[string]int)
    file.File_unzip = make(map[string]*os.File)

    for key, item := range file.File.File {
        file.File_list[string(item.Name)] = key
    }
    return
}

func (f *Zip) Name() string {
    return f.name
}

func (f *Zip) GetFile(name string) (file *zip.File, err error) {
    key, ok := f.File_list[name]
    if ok {
        return f.File.File[key], nil
    }
    return file, errors.New("No file in Zip Archive")
}

// Example:
//  Unzip(name) // default dir = "./tmp/"
//  Unzip(name, "./tmp/")
func (f *Zip) Unzip(name string, params ...string) (file *os.File, err error){
    src_file, err := f.GetFile(name)
    if err != nil {
        return
    }
    file, ok := f.File_unzip[name]
    if ok {
        return
    }
    sr, err := src_file.Open()
    if err != nil {
        return
    }
    defer sr.Close()

    tmp_dir := "./tmp/"
    if len(params) > 0 {
        tmp_dir = params[0]
    }

    file, err = ioutil.TempFile(tmp_dir, "unzipfile")
    if err != nil {
        return
    }
    defer file.Close()

    f.File_unzip[name] = file

    io.Copy(file, sr)
    return
}

func (f *Zip) DeleteUnzip(name string) {
    file, ok := f.File_unzip[name]
    if !ok {
        return
    }
    file_name := file.Name()
    os.Remove(file_name)
    delete(f.File_unzip, name)
}