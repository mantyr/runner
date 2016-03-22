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
    hashfile := GetHashFile(file.Name())
    if hashfile != "e269a4995ad439664251b38951448022706e037b40d243475f1bb3ae74329212" {
        t.Errorf("Error hash256 file, %q", hashfile)
    }
}

func TestRemoveDupRunes(t *testing.T) {
    const nihongo1 = "日 本   ,   語"
    const nihongo2 = " 日,,,    本    ,,,,語     "
    const nihongo3 = " 日 本,語     "
    const nihongo4 = " 日		本	語	"

    val := RemoveDupRunes(nihongo1, " ")
    if val != "日 本 , 語" {
        t.Errorf("Error remove duplicate char, $q", val)
    }

    val = RemoveDupRunes(nihongo1, " ,")
    if val != "日 本 語" {
        t.Errorf("Error remove duplicate char, $q", val)
    }


    val = RemoveDupRunes(nihongo2, " ")
    if val != "日,,, 本 ,,,,語" {
        t.Errorf("Error remove duplicate char, $q", val)
    }
    val = RemoveDupRunes(nihongo2, " ,")
    if val != "日,本 語" {
        t.Errorf("Error remove duplicate char, $q", val)
    }


    val = RemoveDupRunes(nihongo2, " ,", "_")
    if val != "日_本_語" {
        t.Errorf("Error remove duplicate char, $q", val)
    }

    val = RemoveDupRunes(nihongo2, " ,", "_ _")
    if val != "日_ _本_ _語" {
        t.Errorf("Error remove duplicate char, $q", val)
    }

    val = RemoveDupRunes(nihongo3, " ,", " ")
    if val != "日 本 語" {
        t.Errorf("Error remove duplicate char, $q", val)
    }

    val = RemoveDupRunes(nihongo4, " \t", "_")
    if val != "日_本_語" {
        t.Errorf("Error remove duplicate char for tab and other utf8, $q", val)
    }

}

func TestAccessHost(t *testing.T) {
    is := AccessHost("youtube.com", Access_iframe_hosts)
    if !is {
        t.Errorf("Error status in AccessHost, %q", is)
    }

    is = AccessHost("www.youtube.com", Access_iframe_hosts)
    if !is {
        t.Errorf("Error status in AccessHost, %q", is)
    }

    is = AccessHost("sdfkjlaskdjflaskjdf.vk.com", Access_iframe_hosts)
    if !is {
        t.Errorf("Error status in AccessHost, %q", is)
    }

    is = AccessHost("sdfkjlaskdjflaskjdf.vk.com", []string{"test.com", "test2.com"})
    if is {
        t.Errorf("Error status in AccessHost, %q", is)
    }

    is = AccessHost("test2.рф", []string{"test.com", "test2.рф", "vk.com"})
    if !is {
        t.Errorf("Error status in AccessHost, %q", is)
    }
}