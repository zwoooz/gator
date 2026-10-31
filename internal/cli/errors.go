package cli

import "errors"

// ErrUsage is wrapped by errors caused by incorrect command usage
// main uses it to decide whether to show help and exit with code 2
var ErrUsage = errors.New("incorrect usage")
