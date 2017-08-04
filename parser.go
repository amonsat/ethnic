package ethnic

import (
	"sort"
	"strings"

	"github.com/Amonsat/ethnic/data"
	"github.com/Amonsat/ethnic/tools"
)

func ParseFamilyname(familyname string) (prefLangs []string, suffLangs []string) {
	familyname = strings.ToLower(familyname)
	for key, val := range data.Familyname_Prefixes {
		if strings.HasPrefix(familyname, strings.ToLower(key)+" ") {
			prefLangs = append(prefLangs, val...)
		}
	}
	prefLangs = tools.RemoveDuplicates(prefLangs)
	sort.Strings(prefLangs)

	for key, val := range data.Familyname_Suffixes {
		if strings.HasSuffix(familyname, strings.ToLower(key)) {
			suffLangs = append(suffLangs, val...)
		}
	}
	suffLangs = tools.RemoveDuplicates(suffLangs)
	sort.Strings(suffLangs)
	return
}

func ParseNameSuffixes(nameSuffixes string) (langs []string) {
	nameSuffixes = strings.ToLower(nameSuffixes)

	for key, val := range data.NameSuffixes {
		if strings.Contains(nameSuffixes, val) {
			langs = append(langs, val...)
		}
	}
	langs = tools.RemoveDuplicates(langs)
	sort.Strings(langs)
	return
}
