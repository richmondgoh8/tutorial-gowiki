package main
// https://golang.org/doc/articles/wiki/
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
)

type Page struct {
    Title string
    Body  []byte // Page Content
}

func (p *Page) save() error { // Receiver Page, Method Save, No Parameters, Return Error
	fileName := p.Title + ".txt"
	// Creates a txt file
	return ioutil.WriteFile(fileName, p.Body, 0600) //0600 = read write permission for current user
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("edit.html")
    t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1>,<div>%s</div>", p.Title, p.Body)
}

func main() {
	/*p1 := &Page{Title: "Test Page", Body: []byte("This is a sample Page")}
	p1.save()
	p2, _ := loadPage("Test Page")
	fmt.Println(string(p2.Body))*/
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}