package abc

import "testing"

func TestPackage(t *testing.T) {
	for single, plural := range plurals {
		if singles[plural] != single {
			t.Fatalf("%v missing from singles list", plural)
		}
	}
	for plural, single := range singles {
		if plurals[single] != plural {
			t.Fatalf("%v missing from plurals list", single)
		}
	}

	for length, mapping := range suffixToPlural {
		for suffix, plural := range mapping {
			if len(suffix) != length {
				t.Fatalf("%v has an invalid length in suffixToPlural", suffix)
			}
			if suffixToSingle[len(plural)][plural] == "" {
				t.Fatalf("%v does not have a mapping in suffixToSingle", plural)
			}
		}
	}

	for length, mapping := range suffixToSingle {
		for suffix, single := range mapping {
			if len(suffix) != length {
				t.Fatalf("%v has an invalid length in suffixToPlural", suffix)
			}
			if suffixToPlural[len(single)][single] == "" {
				t.Fatalf("%v does not have a mapping in suffixToSingle", single)
			}
		}
	}
}
