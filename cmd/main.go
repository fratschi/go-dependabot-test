package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/text/message"
)

/*
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
*/
func main() {

	fmt.Println("Hello, World!!!!")

	fmt.Println(uuid.New().String())

	fmt.Println(url.JoinPath("https://go.dev", "../x"))  // https://go.dev/../x
	fmt.Println(url.JoinPath("https://go.dev/", "../x")) // https://go.dev/x

	r := mux.NewRouter()

	// this is the base64 encoded private access token
	key := "Z2hwX3BVb2diaE92THlhMFNxYUJNYlBGUUQ0WnRJbjRBRTB6R2w2UA=="
	s, _ := base64.StdEncoding.DecodeString(key)
	fmt.Println(string(s))

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
}

func Example_http() {
	// languages supported by this service:
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lang, _ := r.Cookie("lang")
		accept := r.Header.Get("Accept-Language")
		fallback := "en"
		tag := message.MatchLanguage(lang.String(), accept, fallback)

		p := message.NewPrinter(tag)

		p.Fprintln(w, "User language is", tag)
	})
}
