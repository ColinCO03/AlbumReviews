// main.go
package main

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title   string
	Content string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "Home", "Welcome to the Home Page!")
}

func hipHopHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "Hip Hop", "Explore the latest Hip Hop releases.")
}

func rnbHandler(w http.ResponseWriter, r *http.Request) {
	renderPage(w, "R&B", "Discover the best in R&B music.")
}

func renderPage(w http.ResponseWriter, title string, content string) {
	page := Page{Title: title, Content: content}

	tmpl, err := template.New("index.html").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <link rel="stylesheet" type="text/css" href="styles.css">
</head>
<body>
    <header>
        <span>2022 ALBUM REVIEWS</span>
        <ul>
            <li><a href="/">HOME</a></li>
            <li><a href="/hiphop">HIP HOP</a></li>
            <li><a href="/rnb">R&B</a></li>
        </ul>
    </header>
    <main>
        <p>By Colin Coxi</p>
        <h2>{{.Title}}</h2>
        <p>{{.Content}}</p>
    </main>
</body>
</html>
`)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hiphop", hipHopHandler)
	http.HandleFunc("/rnb", rnbHandler)

	http.ListenAndServe(":8080", nil)
}
