package version

import (
	"fmt"
)

const Version = "v0.0.1"

func LongVersion() string {
	return fmt.Sprintf("TuringML APIs %s", Version)
}
