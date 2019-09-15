package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommit(t *testing.T) {
	var testHash [20]byte

	tests := map[string]Commit{
		"chore: add something\n":               Commit{Category: "chore", Heading: "add something", Hash: testHash},
		"chore(ci): added new CI stuff\n":      Commit{Category: "chore", Scope: "ci", Heading: "added new CI stuff", Hash: testHash},
		"feat: added a new feature\n":          Commit{Category: "feat", Heading: "added a new feature"},
		"fix!: breaking change\n":              Commit{Category: "fix", Breaking: true, Heading: "breaking change"},
		"fix(security)!: breaking\n":           Commit{Category: "fix", Scope: "security", Breaking: true, Heading: "breaking"},
		"fix!!: breaking\n":                    Commit{Category: "fix!", Breaking: true, Heading: "breaking"},
		"fix(security)(stuff): should break\n": Commit{Category: "fix", Scope: "security(stuff)", Heading: "should break"},
		"chore:really close\n":                 Commit{},
		"perf(): nope\n":                       Commit{Category: "perf()", Heading: "nope"},
		"chore(: bad\n":                        Commit{Category: "chore(", Heading: "bad"},
		": nope\n":                             Commit{Heading: "nope"},
		"fix tests\n":                          Commit{},
		"test(full): a heading\n\nbody is here\nit can have multiple lines": Commit{Category: "test", Scope: "full", Heading: "a heading", Body: "body is here\nit can have multiple lines"},
	}

	for test, expected := range tests {
		err := ParseCommit(test, testHash)
		assert.Equal(t, expected, err)
	}
}
