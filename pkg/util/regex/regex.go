package regex

import (
	"regexp"

	"github.com/mingrammer/commonregex"
)

func all() map[string]*regexp.Regexp {
	return map[string]*regexp.Regexp{
		"alphanumeric_w_spaces":    regexp.MustCompile("^[a-zA-Z0-9_ ]*$"),
		"alphanumeric_wout_spaces": regexp.MustCompile("^[a-zA-Z0-9]*$"),
		"username":                 regexp.MustCompile("^[a-z0-9_-]{6,25}$"),
		"email":                    commonregex.LinkRegex,
		"phone_number_w_ext":       commonregex.PhonesWithExtsRegex,
	}
}

func ByName(ruleName string) *regexp.Regexp {
	return all()[ruleName]
}
