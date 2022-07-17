package namepatternenum

type NamePattern string

const (
	lowercase NamePattern = "lowercase"
	upperCase NamePattern = "uppercase"
	title     NamePattern = "title"
)

func (n NamePattern) ToString() string {
	return string(n)
}

func IsNamePatternEnumValue(input string) bool {
	types := []string{lowercase.ToString(), upperCase.ToString(), title.ToString()}

	for _, t := range types {
		if t == input {
			return true
		}
	}

	return false
}
