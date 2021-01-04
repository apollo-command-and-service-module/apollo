package logging

import (
	"os"
)

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}
	//TODO: Do we want os.Exit(1), this will kill the service.
	log := NewConsole(true)
	log.Info().Msg("Error occurred")
	os.Exit(1)
}
