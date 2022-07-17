package namepatternenum

import "strings"

type NamePatternEnum string

const (
	LowerCase NamePatternEnum = "lowercase"
	UpperCase NamePatternEnum = "uppercase"
)

func (n NamePatternEnum) ToString() string {
	return string(n)
}

func NewNamePatternEnum(input string) NamePatternEnum {
	switch strings.ToLower(input) {
	case LowerCase.ToString():
		return LowerCase
	case UpperCase.ToString():
		return UpperCase
	default:
		return ""
	}
}

func IsNamePatternEnumValue(input string) bool {
	return NewNamePatternEnum(input) != ""
}
