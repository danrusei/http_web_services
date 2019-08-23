package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Danr17/http_web_services/api_layered/internal/handlers"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	logger := log.New(os.Stdout, "gcuk ", log.LstdFlags|log.Lshortfile)

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

	srv := &handlers.Server{}

	logger.Println("server starting")
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}

	return err
}
