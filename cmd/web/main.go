package main

import (
	"log"
	"net/http"
)
func main(){
    mux := http.NewServeMux()
    
    fileServer := http.FileServer(http.Dir("./ui/static/"))
    mux.Handle("GET /static/",http.StripPrefix("/static", fileServer))

    mux.HandleFunc("GET /{$}",home)
    mux.HandleFunc("GET /snippet/view/{id}",snippetView)
    mux.HandleFunc("GET /snippet/create",snippetCreate)
    mux.HandleFunc("POST /snippet/create",snippetCreatePost)

    if err := http.ListenAndServe(":4000",mux); err != nil {
        log.Fatal(err)
    }
}

