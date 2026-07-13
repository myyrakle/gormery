package naming

// commonInitialisms is the set of known initialisms, mirroring golint's list.
//
// gorm's pluralization (via github.com/jinzhu/inflection) appends an uppercase
// "S" to an all-caps trailing initialism, so "ProductURL" becomes "ProductURLS"
// and "UserID" becomes "UserIDS". That is not idiomatic Go. FixInitialismPlural
// rewrites the trailing "S" to a lowercase "s" for these known initialisms.
var commonInitialisms = map[string]bool{
	"ACL": true, "API": true, "ASCII": true, "CPU": true, "CSS": true,
	"DNS": true, "EOF": true, "GUID": true, "HTML": true, "HTTP": true,
	"HTTPS": true, "ID": true, "IP": true, "JSON": true, "LHS": true,
	"QPS": true, "RAM": true, "RHS": true, "RPC": true, "SLA": true,
	"SMTP": true, "SQL": true, "SSH": true, "TCP": true, "TLS": true,
	"TTL": true, "UDP": true, "UI": true, "UID": true, "UUID": true,
	"URI": true, "URL": true, "UTF8": true, "VM": true, "XML": true,
	"XMPP": true, "XSRF": true, "XSS": true,
}

// FixInitialismPlural lowercases a trailing plural "S" that gorm/inflection has
// appended to an all-caps initialism, turning "ProductURLS" into "ProductURLs"
// and "UserIDS" into "UserIDs".
//
// It only rewrites when the trailing uppercase run (minus the final "S") is a
// known initialism, so names that pluralize normally are left untouched:
// "Orders" and "Clothes" end in a lowercase "s", and standalone acronyms like
// "CMS" are not in the list ("CM" is not an initialism).
func FixInitialismPlural(name string) string {
	if len(name) < 2 || name[len(name)-1] != 'S' {
		return name
	}

	body := name[:len(name)-1] // drop the trailing plural "S"

	// Walk back over the trailing run of uppercase ASCII letters.
	i := len(body)
	for i > 0 && body[i-1] >= 'A' && body[i-1] <= 'Z' {
		i--
	}

	if commonInitialisms[body[i:]] {
		return body + "s"
	}

	return name
}
