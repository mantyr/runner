package runner

import (
    "testing"
)

func TestTranslite(t *testing.T) {
    text := "New Toyota"
    test := "New_Toyota"

    val := Translite(text)
    if val != test {
        t.Errorf("Error Translite, %q %q %q", text, test, val)
    }
}

func TestTransliteRus(t *testing.T) {
    text := "Привет Мир"
    test := "Privet_Mir"

    val := Translite(text)
    if val != test {
        t.Errorf("Error Translite, %q %q %q", text, test, val)
    }
}

