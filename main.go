package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Page struct { 
    Title string
    Message string
    Thing string
}

func main(){
    http.HandleFunc("/", index)
    http.HandleFunc("/request", request)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func request(w http.ResponseWriter, r *http.Request){
    t, err := getTemplate("request/index")
    if err != nil {
        http.Error(w, "404 not found", http.StatusNotFound);
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    t.Execute(w, Page{})
}

func index(w http.ResponseWriter, r *http.Request) {
    t, err := getTemplate("index")

    if err != nil {
        http.Error(w, "404 not found", http.StatusNotFound);
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }

    t.Execute(w, Page{})
}

func getTemplate(name string) (*template.Template, error){
    x := "static/" + name + ".html"
    t, err := template.ParseFiles(x)

    if err != nil {
        return nil, err
    }
    return t, nil
}
