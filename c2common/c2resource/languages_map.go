package c2resource

func getLanguages() map[byte]string {
	return map[byte]string{
		0x00: "CHI", // (Chinese)
		0x01: "JPN", // (Japanese)
		0x02: "KOR", // (Korean)
		0x03: "ENG", // (English)
		0x04: "DEU", // (German)
		0x05: "ESP", // (Spanish)
		0x06: "POR", // (Portuguese)
		0x07: "ITA", // (Italian)
		0x08: "POL", // (Polish)
		0x09: "RUS", // (Russian)
		0x0A: "SIN", //
	}
}

// GetLanguageLiteral returns string representation of language code
func GetLanguageLiteral(code byte) string {
	languages := getLanguages()
	return languages[code]
}

func getCharsets() map[string]string {
	return map[string]string{
		"CHI": "CHI",    // (Chinese)
		"JPN": "JPN",    // (Japanese)
		"KOR": "KOR",    // (Korean)
		"ENG": "LATIN",  // (English)
		"DEU": "LATIN",  // (German)
		"FRA": "LATIN",  // (French)
		"ESP": "LATIN",  // (Spanish)
		"POR": "LATIN",  // (Portuguese)
		"ITA": "LATIN",  // (Italian)
		"POL": "LATIN2", // (Polish)
		"RUS": "CYR",    // (Russian)
		"SIN": "LATIN",  //
	}
}

// GetFontCharset returns string representation of font charset
func GetFontCharset(language string) string {
	charset := getCharsets()
	return charset[language]
}

// GetLabelModifier returns modifier for language
/* modifiers for labels (used in string tables)
modifier is something like that:
english table:       polish table:
key  | value         key  |  value
#1   | v1                 |
#4   | v2            #4   | v1
#5   | v3            #5   | v2
#8   | v4            #8   | v3
So, GetLabelModifier returns value of offset in locale languages table
*/
// some of values need to be set up. For now values with "checked" comment
// was tested and works fine.
func GetLabelModifier(language string) int {
	modifiers := map[string]int{
		"CHI": 0, // (Chinese)
		"JPN": 0, // (Japanese)
		"KOR": 0, // (Korean)
		"ENG": 0, // (English) // checked
		"DEU": 0, // (German) // checked
		"FRA": 0, // (French)
		"ESP": 0, // (Spanish)
		"POR": 0, // (Portuguese)
		"ITA": 0, // (Italian) // checked
		"POL": 1, // (Polish) // checked
		"RUS": 0, // (Russian)
		"SIN": 0, //
	}

	return modifiers[language]
}
