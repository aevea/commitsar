package text

import (
	"github.com/outillage/quoad"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommit(t *testing.T) {
	var testHash [20]byte

	tests := map[string]quoad.Commit{
		"chore: add something\n":               quoad.Commit{Category: "chore", Heading: "add something", Hash: testHash},
		"chore(ci): added new CI stuff\n":      quoad.Commit{Category: "chore", Scope: "ci", Heading: "added new CI stuff", Hash: testHash},
		"feat: added a new feature\n":          quoad.Commit{Category: "feat", Heading: "added a new feature"},
		"fix!: breaking change\n":              quoad.Commit{Category: "fix", Breaking: true, Heading: "breaking change"},
		"fix(security)!: breaking\n":           quoad.Commit{Category: "fix", Scope: "security", Breaking: true, Heading: "breaking"},
		"fix!!: breaking\n":                    quoad.Commit{Category: "fix!", Breaking: true, Heading: "breaking"},
		"fix(security)(stuff): should break\n": quoad.Commit{Category: "fix", Scope: "security(stuff)", Heading: "should break"},
		"chore:really close\n":                 quoad.Commit{},
		"perf(): nope\n":                       quoad.Commit{Category: "perf()", Heading: "nope"},
		"chore(: bad\n":                        quoad.Commit{Category: "chore(", Heading: "bad"},
		": nope\n":                             quoad.Commit{Heading: "nope"},
		"fix tests\n":                          quoad.Commit{},
		"test(full): a heading\n\nbody is here\nit can have multiple lines": quoad.Commit{Category: "test", Scope: "full", Heading: "a heading", Body: "body is here\nit can have multiple lines"},
	}

	for test, expected := range tests {
		err := ParseCommit(test, testHash)
		assert.Equal(t, expected, err)
	}
}
