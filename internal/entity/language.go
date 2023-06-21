package entity

import "regexp"

const (
	languageCodeRegex   = "^[a-z]+$"
	languageCodeSize    = 2
	defaultLanguageCode = "en"
)

func ValidatedLanguage(rawLanguage *string) string {
	if rawLanguage == nil {
		return defaultLanguageCode
	}

	validLanguage, err := regexp.MatchString(languageCodeRegex, *rawLanguage)
	if len(*rawLanguage) != languageCodeSize || err != nil || !validLanguage {
		return defaultLanguageCode
	}

	return *rawLanguage
}
