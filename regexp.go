package validate

import (
	"regexp"
)

type lazyRegexp struct {
	p string
	r *regexp.Regexp
}

func (l *lazyRegexp) MatchString(str string) bool {
	if l.r == nil {
		l.r = regexp.MustCompile(l.p)
	}

	return l.r.MatchString(str)
}

var rxAlpha = lazyRegexp{p: `(?i)^\p{L}+$`}
var rxAlphaLatin = lazyRegexp{p: `(?i)^[a-z]+$`}
