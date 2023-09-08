package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/go-fed/activitypub"
  "github.com/unmsmfisi-socialapplication/social_app/internal/multipublication/application"
)

func main() {
  // Load the multipublication configuration.
  config, err := config.LoadMultipublicationConfig()
  if err != nil {
    log.Fatal(err)
  }

  // Create a new HTTP server.
  server := http.Server{
    Addr: ":8080",
  }

  // Create the multipublication handler.
  handler := application.NewMultipublicationHandler(config)

  // Register the multipublication handler.
  http.HandleFunc("/multipublication", handler.Handle)

  // Start the HTTP server.
  err = server.ListenAndServe()
  if err != nil {
    log.Fatal(err)
  }
}
