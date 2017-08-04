package ethnic

import (
	"sort"
	"strings"

	"github.com/Amonsat/ethnic/data"
)

func ParseFamilyname(familyname string) (prefLangs []string, suffLangs []string) {
	familyname = strings.ToLower(familyname)
	for key, val := range data.Prefixes {
		if strings.HasPrefix(familyname, strings.ToLower(key)+" ") {
			prefLangs = append(prefLangs, val...)
		}
	}
	prefLangs = removeDuplicates(prefLangs)
	sort.Strings(prefLangs)

	for key, val := range data.Suffixes {
		if strings.HasSuffix(familyname, strings.ToLower(key)) {
			suffLangs = append(suffLangs, val...)
		}
	}
	suffLangs = removeDuplicates(suffLangs)
	sort.Strings(suffLangs)
	return
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}
