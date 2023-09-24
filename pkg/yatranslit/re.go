package yatranslit

import (
	"fmt"
	"regexp"
)

const dash = "-"

type compiledRegexps struct {
	alpha, doubleDash *regexp.Regexp
}

// Premain compilation and cache regexp
var re = compileRE()

func compileRE() (result compiledRegexps) {
	result.alpha = regexp.MustCompile(fmt.Sprintf(`[^\w\%s]+`, dash))
	result.doubleDash = regexp.MustCompile(fmt.Sprintf(`%s%s+`, dash, dash))

	return result
}
