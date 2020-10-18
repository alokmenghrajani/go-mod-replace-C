package go_mod_replace_C

import (
	"fmt"
	x "github.com/alokmenghrajani/go-mod-replace-A"
)

func LibC() string {
	return fmt.Sprintf("LibC: %s, %s\n", x.LibAanotherFunc(), x.LibA())
}
