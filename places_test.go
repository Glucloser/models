package models

import "testing"

func TestSearchNameFor(t *testing.T) {
	simpleName := "Cheers"
	expectedSimpleSearchName := "cheers"
	simpleSearchName := searchNameFor(simpleName)
	if expectedSimpleSearchName != simpleSearchName {
		t.Fatalf("%s should equal %s", simpleSearchName, expectedSimpleSearchName)
	}

	twoWordName := "Central Perk"
	expTwoWordSearchName := "central_perk"
	twoWordSearchName := searchNameFor(twoWordName)
	if expTwoWordSearchName != twoWordSearchName {
		t.Fatalf("%s should equal %s", twoWordSearchName, expTwoWordSearchName)
	}

	punctuationName := "Bob's Burgers"
	expPunctuationSearchName := "bob$s_burgers"
	punctuationSearchName := searchNameFor(punctuationName)
	if expPunctuationSearchName != punctuationSearchName {
		t.Fatalf("%s should equal %s", punctuationSearchName, expPunctuationSearchName)
	}

	unicodeName := "Krüsty Burger"
	expUnicodeSeachName := "kru$sty_burger"
	unicodeSearchName := searchNameFor(unicodeName)
	if unicodeSearchName != expUnicodeSeachName {
		t.Fatalf("%s should equal %s", unicodeSearchName, expUnicodeSeachName)
	}
}
