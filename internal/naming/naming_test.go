package naming

import "testing"

func TestFixInitialismPlural(t *testing.T) {
	cases := map[string]string{
		// gorm/inflection appends an uppercase "S" to trailing initialisms.
		"ProductURLS": "ProductURLs",
		"ImageURLS":   "ImageURLs",
		"URLS":        "URLs",
		"UserIDS":     "UserIDs",
		"IDS":         "IDs",
		"APIS":        "APIs",
		"UUIDS":       "UUIDs",

		// Normal plurals end in a lowercase "s" and must stay untouched.
		"Orders":  "Orders",
		"Clothes": "Clothes",
		"People":  "People",

		// Trailing run that is not a known initialism must stay untouched.
		"CMS": "CMS", // "CM" is not an initialism
		"AS":  "AS",  // "A" is not an initialism

		// Degenerate inputs.
		"S": "S",
		"":  "",
	}

	for in, want := range cases {
		if got := FixInitialismPlural(in); got != want {
			t.Errorf("FixInitialismPlural(%q) = %q, want %q", in, got, want)
		}
	}
}
