package main

import (
	"html/template"
	"log"
	"net/http"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeList(writer http.ResponseWriter, list []string) {
	// YOUR CODE HERE: Use the template library to parse the contents
	// of the "list.html" file. You'll get a template value and an
	// error value.
	html, err := template.ParseFiles("list.html")
	// Pass the error value to the "check" function.
	check(err)
	// Now call the "Execute" method on the template. It should write
	// its output to the "writer" parameter, and it should use the
	// "list" parameter as its data value.
	err = html.Execute(writer, list)
	// You'll get another error value back from "Execute", which
	// should be passed to "check".
	check(err)
}

func fruitHandler(writer http.ResponseWriter, request *http.Request) {
	writeList(writer, []string{"apples", "oranges", "pears"})
}

func meatHandler(writer http.ResponseWriter, request *http.Request) {
	writeList(writer, []string{"chicken", "beef", "lamb"})
}

func main() {
	http.HandleFunc("/fruit", fruitHandler)
	http.HandleFunc("/meat", meatHandler)
	// 8080 busy, so 8010 port is used to listen and server
	err := http.ListenAndServe("localhost:8010", nil)
	log.Fatal(err)
}
