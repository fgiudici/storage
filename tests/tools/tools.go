// +build tools

package tools

// Importing the packages here will allow to vendor those via
// `go mod vendor`.

import (
	_ "github.com/cpuguy83/go-md2man"
	_ "github.com/pquerna/ffjson"
	_ "github.com/vbatts/git-validation"
)
