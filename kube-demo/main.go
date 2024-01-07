package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

func main() {
	started := time.Now()

	if err := connectDatabase(); err != nil {
		log.Fatal(err)
	}
	defer databaseConn.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Now().Sub(started)

		if duration.Seconds() > 1000 {
			slog.Info("Timeout triggered")
			w.WriteHeader(http.StatusInternalServerError)
			_, err := w.Write([]byte(`I'm timed out'`))
			if err != nil {
				slog.Error(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write([]byte(`Hello gopher`))
			if err != nil {
				slog.Error(err.Error())
			}
		}
	})

	http.HandleFunc("/alligator", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello Mr Alligator")
		if err != nil {
			slog.Error(err.Error())
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
