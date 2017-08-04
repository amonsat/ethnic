package ethnic

import (
	"strings"
)

func ParseFamilyname(familyname string) (prefLangs []string, suffLangs []string) {
	familyname = strings.ToLower(familyname)
	for key, val := range prefixes {
		if strings.HasPrefix(familyname, strings.ToLower(key)+" ") {
			prefLangs = val
		}
	}

	for key, val := range suffixes {
		if strings.HasPrefix(familyname, strings.ToLower(key)+" ") {
			suffLangs = val
		}
	}
	return
}
