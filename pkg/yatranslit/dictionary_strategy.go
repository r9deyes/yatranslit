package yatranslit

import (
	"strings"
)

type Translater interface {
	Translate(text string, maxLength int) string
}

type DictionaryStrategy struct {
	dictionary TranslitDict
}

func NewDictionaryStrategy() *DictionaryStrategy {
	return &DictionaryStrategy{
		dictionary: Dictionary,
	}
}

func (ds *DictionaryStrategy) SetDictionary(dictionary TranslitDict) {
	ds.dictionary = dictionary
}

func (ds *DictionaryStrategy) Translate(text string, maxLength int) (result string) {
	result = ds.sanitize(ds.Parse(text))
	if maxLength > 0 && maxLength < len(result) {
		return result[0:maxLength]
	}
	return result
}

func (ds *DictionaryStrategy) Parse(text string) string {
	var (
		previous, current, next rune
		result                  []rune
	)

	lowered := []rune(strings.ToLower(text))
	textLen := len(lowered)
	safeIndex := func(idx int) rune {
		if idx >= 0 && idx < textLen {
			return lowered[idx]
		}
		return 0
	}
	// result len can be greater
	result = make([]rune, 0, textLen)

	previous = safeIndex(0)
	for i := 0; i < textLen; i++ {
		current = safeIndex(i)
		next = safeIndex(i + 1)
		result = append(result, ds.translateLetter(current, previous, next)...)
		previous = current
	}

	return string(result)
}

func (ds *DictionaryStrategy) translateLetter(current rune, previous rune, next rune) []rune {
	if subMap, ok := ds.dictionary[string(current)]; ok {
		if len(subMap) == 1 {
			// FIXME getting only one existing value
			for _, v := range subMap {
				return []rune(v)
			}
		} else {
			return ds.handleSpecialCases(current, previous, next, subMap)
		}
	}

	return []rune{current}
}

// handleSpecialCases: Special cases where dictionary array contains sub arrays of additional rules
func (ds *DictionaryStrategy) handleSpecialCases(current rune, previous rune, next rune, subMap map[string]string) []rune {
	var (
		prevRS, nextRS, combination []rune
	)
	combination = make([]rune, 2)

	if previous != 0 {
		combination[0], combination[1] = previous, current

		if val, ok := subMap[string(combination)]; ok {
			return []rune(val)
		}
		prevRS = []rune{previous}
	}

	if next != 0 {
		combination[0], combination[1] = current, next

		if val, ok := subMap[string(combination)]; ok {
			return []rune(val)
		}
		nextRS = []rune{next}
	}

	// prev + current + next with trimming empty
	combination = append(append(prevRS, current), nextRS...)

	if val, ok := subMap[string(combination)]; ok {
		return []rune(val)
	}

	if val, ok := subMap[anyOther]; ok {
		return []rune(val)
	}
	return []rune{current}
}

func (ds *DictionaryStrategy) sanitize(parsed string) string {
	noWhiteSpaces := strings.Replace(parsed, " ", dash, -1)

	// Remove all characters but words, numbers and dashes `[^\w\-]+`
	alphaNum := re.alpha.ReplaceAllLiteralString(noWhiteSpaces, "")

	//Remove double dashes `--+`
	noDoubleDashed := re.doubleDash.ReplaceAllLiteralString(alphaNum, dash)

	return strings.Trim(noDoubleDashed, dash)
}
