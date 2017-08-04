package ethnic

import (
	"sort"
	"strings"

	"github.com/Amonsat/ethnic/data"
	"github.com/Amonsat/ethnic/tools"
)

func ParseFamilyname(familyname string) (prefLangs []string, suffLangs []string) {
	familyname = strings.ToLower(familyname)
	for key, val := range data.Prefixes {
		if strings.HasPrefix(familyname, strings.ToLower(key)+" ") {
			prefLangs = append(prefLangs, val...)
		}
	}
	prefLangs = tools.RemoveDuplicates(prefLangs)
	sort.Strings(prefLangs)

	for key, val := range data.Suffixes {
		if strings.HasSuffix(familyname, strings.ToLower(key)) {
			suffLangs = append(suffLangs, val...)
		}
	}
	suffLangs = tools.RemoveDuplicates(suffLangs)
	sort.Strings(suffLangs)
	return
}
