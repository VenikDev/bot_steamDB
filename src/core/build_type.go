package core

import (
	"github.com/samber/mo"
	"os"
)

const (
	DEBUG   = "debug"
	RELEASE = "release"
)

// BuildIsDebug
// This code defines a function named `BuildIsDebug()` that returns an `Option` (
// a value that may or may not be present) of `bool` type.
// The function first gets the value of the `BUILD_TYPE` environment variable using `os.Getenv(
// )` method and stores it in `buildType` variable.
// If the value of `buildType` is equal to a constant `DEBUG`,
// the function returns `Some(true)` (returns value `true` wrapped in `Some`). Else,
// if the value of `buildType` is equal to a constant `RELEASE`,
// the function returns `Some(false)` (returns value `false` wrapped in `Some`).
// If none of the above conditions are met, the function returns `None` (no value is present).
// (!!!) Using mo.Option for test only
func BuildIsDebug() mo.Option[bool] {
	buildType := os.Getenv("BUILD_TYPE")

	if buildType == DEBUG {
		return mo.Some(true)
	} else if buildType == RELEASE {
		return mo.Some(false)
	}

	return mo.None[bool]()
}
