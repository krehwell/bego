package main

import (
    "fmt"
    "unicode/utf8"
)

func examineRune(r rune) {
    if r == 't' {
        fmt.Println("found tee")
    } else if r == 'ส' {
        fmt.Println("found so sua")
    }
}

func StringsAndRunes() {
    const s = "สวัสดี" // special runes might have different width each

    fmt.Println("length of thai word", len(s))

    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    fmt.Println("Run Count", utf8.RuneCountInString(s))

    for idx, char := range s {
        fmt.Printf("%c starts at %d\n", char, idx)
    }

    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d with width %d\n", runeValue, i, width)
        w = width

        examineRune(runeValue)
    }
}

