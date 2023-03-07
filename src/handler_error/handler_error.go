package herr

import "bot_steamDB/src/clog"

// HandlerError
// This is a Go function named "HandlerError" that takes two parameters: "err" and "msg".
// The function checks if the "err" parameter is not nil (meaning there is an error). If there is an error,
// the function logs the "msg" and the "Error" string (to indicate it's an error),
// as well as the error itself using a logger called "clog.Logger", which is assumed to have a method called "Error".
func HandlerError(err error, msg string) {
	if err != nil {
		clog.Logger.Error(msg, "Error", err)
	}
}

// HandlerFatal
// The code defines a function named HandlerFatal that takes in two parameters - an error and a string message.
// It checks if the error is not nil, and if it is not,
// the function logs a fatal error message using a logger named clog.
// Logger along with the given message and the error object. If the error is nil,
// the function returns without doing anything.
func HandlerFatal(err error, msg string) {
	if err != nil {
		clog.Logger.Fatal(msg, "Error", err)
	}
}
