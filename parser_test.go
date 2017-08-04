package ethnic

import (
	"reflect"
	"testing"
)

func TestParseFamilyname(t *testing.T) {

	tests := []struct {
		name          string
		familyname    string
		wantPrefLangs []string
		wantSuffLangs []string
	}{
		// TODO: Add test cases.
		{"ov test", "Lomonosov", []string{}, []string{"Russian", "all Eastern Slavic languages"}},
		{"son test", "Johnson", []string{}, []string{"English", "French", "German", "Icelandic", "Norwegian", "Swedish"}},
		{"de la test", "de la Vega", []string{"French", "Italian", "Portuguese", "Spanish"}, []string{"European", "Kurdish", "except French Frisian"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrefLangs, gotSuffLangs := ParseFamilyname(tt.familyname)
			if !reflect.DeepEqual(gotPrefLangs, tt.wantPrefLangs) {
				t.Errorf("ParseFamilyname() gotPrefLangs = %v, want %v", gotPrefLangs, tt.wantPrefLangs)
			}
			if !reflect.DeepEqual(gotSuffLangs, tt.wantSuffLangs) {
				t.Errorf("ParseFamilyname() gotSuffLangs = %v, want %v", gotSuffLangs, tt.wantSuffLangs)
			}
		})
	}
}
