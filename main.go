package main

import (
	"html/template"
	"net/http" 

	"github.com/gorilla/mux"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseGlob("template/*.html"))
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome <h1>to Go Web</h1>")
	temp.ExecuteTemplate(w, "index.html", nil)
}

func handleHipHop(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome <h1>to Go Web</h1>")
	temp.ExecuteTemplate(w, "hiphop.html", nil)

}

func handleRNB(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Welcome <h1>to Go Web</h1>")
	temp.ExecuteTemplate(w, "rnb.html", nil)
}
func main() {
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	http.HandleFunc("/", handleFunc)
	http.HandleFunc("/home", )

}
