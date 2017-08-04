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

func TestParseNameSuffixes(t *testing.T) {
	tests := []struct {
		name         string
		nameSuffixes string
		wantLangs    []string
	}{
		{"base test", "md, III", []string{"American", "British"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLangs := ParseNameSuffixes(tt.nameSuffixes); !reflect.DeepEqual(gotLangs, tt.wantLangs) {
				t.Errorf("ParseNameSuffixes() = %v, want %v", gotLangs, tt.wantLangs)
			}
		})
	}
}
