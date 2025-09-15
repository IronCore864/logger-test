package main

import (
	"fmt"
	"os"

	"github.com/ironcore864/logger-test/logger"
)

func main() {
	myLogger := logger.New(os.Stdout, "MyApp: ")
	logger.SetLogger(myLogger)

	logger.Noticef("Starting application...")
	logger.Noticef("This is a formatted log: %d", 42)
	logger.Noticef("Application finished.")

	myApp := "hello world"
	logger.Noticef("%s", fmt.Sprintf("Starting application: %s ...", myApp))

}
