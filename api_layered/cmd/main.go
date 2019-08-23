package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	//	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

	/* Start Database
		db, err := database.Setup()
		if err != nil {
			return errors.Wrap(err, "setup database")
		}
		defer db.Close()

		srv := &handlersserver{
	      db: db,
	  }

	*/

	return nil
}
