package rss

import (
	"database/sql"
	"strings"
	"time"
)

// ToNullString converts a string to a sql.NullString, valid only if non-empty.
func ToNullString(str string) sql.NullString {
	trimmed := strings.TrimSpace(str)
	return sql.NullString{
		String: str,
		Valid:  trimmed != "",
	}
}

var timeFormats = []string{
	time.RFC1123,
	time.RFC1123Z,
	time.RFC822,
	time.RFC822Z,
}

// ToNullTime parses a feed date, trying common RSS layouts in order. It returns
// an invalid sql.NullTime if the string is empty or matches none of them.
func ToNullTime(str string) sql.NullTime {
	str = strings.TrimSpace(str)
	if str == "" {
		return sql.NullTime{}
	}

	for _, format := range timeFormats {
		parsed, err := time.Parse(format, str)
		if err != nil {
			continue
		}
		return sql.NullTime{
			Time:  parsed,
			Valid: true,
		}
	}

	return sql.NullTime{}
}
