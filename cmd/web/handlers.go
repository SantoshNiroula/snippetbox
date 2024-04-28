package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


func home(w http.ResponseWriter, r *http.Request){
    w.Header().Add("Server","Go")

    files := []string {
        "./ui/html/base.tmpl",
        "./ui/html/partials/nav.tmpl",
        "./ui/html/pages/home.tmpl",
    }

    ts, err := template.ParseFiles(files...)
    if err != nil {
        log.Printf(err.Error())
        http.Error(w,"Internal Server Error", http.StatusInternalServerError)
        return
    }

    err = ts.ExecuteTemplate(w,"base",nil)
    if err != nil {
        log.Printf(err.Error())
        http.Error(w,"Internal Server Error", http.StatusInternalServerError)
    }

}
func snippetView(w http.ResponseWriter, r *http.Request){
    id, err := strconv.Atoi(r.PathValue("id"))
    if err != nil || id < 1 {
        http.NotFound(w,r)
        return
    }
    msg := fmt.Sprintf("Will dispaly snippet with ID %d",id)
    w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Snippet Create"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request){
    w.Header().Add("Server", "Go")
    w.Header().Add("Debug", "True")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Snippet saved"))
}





