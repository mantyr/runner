package runner

import (
    "testing"
)

func TestGetFileExt(t *testing.T) {
    ext := GetFileExt("/dir/dir2/file.tag.gz")
    if ext != "gz" {
        t.Errorf("Error get ext file, %q", ext)
    }
    ext = GetFileExt("/dir/dir2/file.tar")
    if ext != "tar" {
        t.Errorf("Error get ext file, %q", ext)
    }
    ext = GetFileExt("/dir/dir2/file")
    if ext != "" {
        t.Errorf("Error get ext file, %q", ext)
    }
}
