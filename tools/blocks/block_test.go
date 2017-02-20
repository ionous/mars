package blocks

import (
	"encoding/json"
	. "github.com/ionous/mars/script"
	"github.com/ionous/mars/std"
	"github.com/ionous/mars/tools/inspect"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCallbackUnknown tests compiler failure when an action does not exist
func TestCallbackUnknown(t *testing.T) {
	what := The("cabinet", IsKnownAs("armoire"))
	if types, err := inspect.NewTypes(std.Std()); assert.NoError(t, err, "types") {
		if db, err := NewModelMaker(types).Compute(what); assert.NoError(t, err, "model") {
			if prettyBytes, err := json.MarshalIndent(db, "", " "); assert.NoError(t, err, "model") {
				t.Log(string(prettyBytes))
			}

			// NOW: generate blocks; then, generate text.

		}
	}

}
