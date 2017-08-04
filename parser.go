package ethnic

import (
	"strings"

	"./data"
)

func ParseFamilyname(familyname string) (prefLangs []string, suffLangs []string) {
	familyname = strings.ToLower(familyname)
	for key, val := range data.Prefixes {
		if strings.HasPrefix(familyname, strings.ToLower(key)+" ") {
			prefLangs = val
		}
	}

	for key, val := range data.Suffixes {
		if strings.HasPrefix(familyname, strings.ToLower(key)+" ") {
			suffLangs = val
		}
	}
	return
}
