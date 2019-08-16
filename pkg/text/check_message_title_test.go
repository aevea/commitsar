package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMessageTitle(t *testing.T) {
	tests := map[string]error{
		"chore: add something":               nil,
		"chore(ci): added new CI stuff":      nil,
		"feat: added a new feature":          nil,
		"fix!: breaking change":              nil,
		"fix(security)!: breaking":           nil,
		"fix!!: breaking":                    errTitleNonConform,
		"fix(security)(stuff): should break": errTitleNonConform,
		"chore:really close":                 errTitleNonConform,
		"perf(): nope":                       errTitleNonConform,
		"chore(: bad":                        errTitleNonConform,
		": nope":                             errTitleNonConform,
		"fix tests":                          errTitleNonConform,
	}

	for test, expected := range tests {
		err := CheckMessageTitle(test)
		assert.Equal(t, expected, err)
	}
}
