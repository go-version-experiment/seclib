package seclib

import (
	"fmt"

	"github.com/go-version-experiment/baselib"
)

const Version = "0.1.0"

//go:noinline
func GetVersion() string {
	return fmt.Sprintf("%s w/ baselib %s", Version, baselib.GetVersion())
}
