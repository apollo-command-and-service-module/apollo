package logging

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}
	log := NewConsole(true)
	log.PrintErrorf("An error occurred: %s", err)
	//TODO: Do we want os.Exit(1), this will kill the service.
	//Exit is used when we need to abort the program immediately so not sure if we want that each time an error is thrown.
	//os.Exit(1)
}
