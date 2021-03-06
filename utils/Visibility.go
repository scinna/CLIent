package utils

import "strings"

type Visibility int

const (
	VisibilityPublic   Visibility = 0
	VisibilityUnlisted Visibility = 1
	VisibilityPrivate  Visibility = 2
)

func (v Visibility) IsValid() bool {
	return v >= 0 && v <= 2
}

func (v Visibility) ToString() string {
	switch v {
	case 0:
		return "public"
	case 1:
		return "unlisted"
	case 2:
		return "private"
	}

	return ""
}

func VisibilityFromString(v string) Visibility {
	switch strings.ToLower(v) {
	case "public":
		return VisibilityPublic
	case "unlisted":
		return VisibilityUnlisted
	case "private":
		return VisibilityPrivate
	}

	return -1
}
