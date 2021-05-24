//Package an prefixes indefinite articles onto a string (a or an).
package an

import (
	"strings"
)

func startsWithVowel(s string) bool {
	switch s[0] {
	case 'a', 'e', 'i', 'o', 'u',
		'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func acronymException(s string) bool {
	switch s[0] {
	case 'U', 'F', 'H', 'L', 'M', 'N', 'R', 'S', 'X':
		return true
	default:
		return false
	}
}

func isException(s string) bool {
	switch s {
	case
		// Nouns: eu like y
		"eunuch",
		"eucalyptus",
		"eugenics",
		"eulogy",
		"euphemism",
		"euphony",
		"euphoria",
		"eureka",

		// Adjectives: eu like y
		"euro", "european", "euphemistic", "euphonic", "euphoric",

		// Adverbs: eu like y
		"euphemistically", "euphonically", "euphorically",

		// Nouns: silent h
		"heir", "heir's", "heiress", "heiresses", "herb", "homage", "honesty", "honor", "honors", "honour", "honored", "honoured", "hour", "hours",

		// Adjectives: silent h
		"honest", "honorous",

		// Adverbs: silent h
		"honestly", "hourly",

		// Nouns: o like w
		"one", "ouija",

		// Adjectives: o like w
		"once",

		// Adverbs: o like w

		// Nouns: u like y
		"ubiquity", "udometer", "ufo", "uke", "ukelele", "ululate", "unicorn", "unicorn's", "unicycle", "uniform",
		"unify", "union", "unions", "unison", "unit", "units", "unity", "universe", "university", "university's", "upas", "ural", "uranium",
		"urea", "ureter", "urethra", "urine", "urologist", "urologist's", "urology", "urus", "usage", "use", "user", "usual", "usurp",
		"usury", "utensil", "uterus", "utility", "utopia", "utricle", "uvarovite", "uvea", "uvula", "utah", "utahn",

		// Adjectives: u like y
		"ubiquitous", "ugandan", "ukrainian", "unanimous", "unicameral", "unified", "unique", "unisex",
		"universal", "urinal", "urological", "useful", "useless", "usurious", "usurped", "utilitarian",
		"utopic",

		// Adverbs: u like y
		"ubiquitously", "unanimously", "unicamerally", "uniquely", "universally", "urologically", "usefully", "uselessly", "usuriously",

		// Nouns: y like i
		"yttria", "yggdrasil", "ylem", "yperite", "ytterbia", "ytterbium", "yttrium",

		// Adjectives: y like i
		"ytterbous", "ytterbic", "yttric",

		// Single letters
		"f", "h", "l", "m", "n", "r", "s", "u", "x":

		return true
	default:
		return false
	}
}

//A returns the given string prefixed with the sensible
//indefinite article. Either "a ", or "an ".
func A(given string) string {
	if given == "" {
		return ""
	}

	word := strings.SplitN(given, " ", 2)[0]
	word = strings.SplitN(word, "-", 2)[0]

	//acronyms
	if len(word) > 1 && word[0] <= 'Z' && word[1] <= 'Z' {
		if startsWithVowel(word) == acronymException(word) {
			return "a " + given
		}
		return "an " + given
	}

	word = strings.ToLower(word)

	if startsWithVowel(word) != isException(word) {
		return "an " + given
	}
	return "a " + given
}
