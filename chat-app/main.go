package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

// TemplateHandler templ represents a single template
type TemplateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, req)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags

	r := newRoom()

	http.Handle("/", &TemplateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	log.Printf("Starting web server on %v", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
