package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var f = fmt.Printf

func main() {
	f("To Upper: %s\n", strings.ToUpper("Hello THERE"))
	f("To Lower: %s\n", strings.ToLower("Hello THERE"))

	caser := cases.Title(language.English)
	f("%s\n", caser.String("tHis will be A title"))

	f("EqualFold: %v\n", strings.EqualFold("Blessed", "BLESsed"))
	f("EqualFold: %v\n", strings.EqualFold("Blessed", "BLESse"))

	f("Index: %v\n", strings.Index("Blessed", "se"))
	f("Index: %v\n", strings.Index("Blessed", "Se"))

	f("Count: %v\n", strings.Count("Blessed", "e"))
	f("Count: %v\n", strings.Count("Blessed", "E"))
	f("Repeat: %s\n", strings.Repeat("ab", 5))

	f("TrimSpace: %s\n", strings.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s\n", strings.TrimLeft(" \tThis is a line. \n", "\n\t "))
	f("TrimRight: %s\n", strings.TrimRight(" \tThis is a\t line. \n", "\n\t "))

	f("Prefix: %v\n", strings.HasPrefix("Blessed", "Bl"))
	f("Prefix: %v\n", strings.HasPrefix("Blessed", "bl"))
	f("Suffix: %v\n", strings.HasSuffix("Blessed", "ed"))
	f("Suffix: %v\n", strings.HasSuffix("Blessed", "ED"))

	t := strings.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = strings.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))

	f("%s\n", strings.Split("abcd efg", ""))
	f("%s\n", strings.Replace("abcd efg", "", "_", -1))
	f("%s\n", strings.Replace("abcd efg", "", "_", 4))
	f("%s\n", strings.Replace("abcd efg", "", "_", 2))

	f("SplitAfter: %s\n", strings.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", strings.TrimFunc("123 abc ABC \t .", trimFunction))
}
