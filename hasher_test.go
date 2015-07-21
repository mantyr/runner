package runner

import (
    "testing"
    "os"
)

func TestGetHashFile(t *testing.T) {
    file, err := os.Open("./testdata/test.jpg")
    if err != nil {
        t.Errorf("Error open file, %q", err)
    }
    hashfile := GetHashFile(file)
    if hashfile != "e269a4995ad439664251b38951448022706e037b40d243475f1bb3ae74329212" {
        t.Errorf("Error hash256 file, %q", hashfile)
    }
}
