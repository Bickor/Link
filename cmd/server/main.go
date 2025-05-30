package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

type Person struct {
	Name    string
	MetAt   string
	WorksAt string
	Fact    string
}

var (
	people []Person
	// templates = template.Must(template.ParseGlob("templates/*.html"))
	templates = loadTemplates()
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/add", handleAdd)
	http.HandleFunc("/edit", handleEditForm)
	http.HandleFunc("/update", handleUpdate)
	http.HandleFunc("/delete", handleDelete)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	p := Person{
		Name:    r.FormValue("name"),
		MetAt:   r.FormValue("met_at"),
		WorksAt: r.FormValue("works_at"),
		Fact:    r.FormValue("fact"),
	}

	people = append(people, p)

	err := templates.ExecuteTemplate(w, "person_row.html", struct {
		Index  int
		Person Person
	}{Index: len(people) - 1, Person: p})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleEditForm(w http.ResponseWriter, r *http.Request) {
	idx, _ := strconv.Atoi(r.FormValue("index"))
	if idx < 0 || idx >= len(people) {
		http.Error(w, "Index out of bounds", http.StatusBadRequest)
		return
	}

	err := templates.ExecuteTemplate(w, "person_edit_form.html", struct {
		Index  int
		Person Person
	}{Index: idx, Person: people[idx]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
	idx, _ := strconv.Atoi(r.FormValue("index"))
	if idx < 0 || idx >= len(people) {
		http.Error(w, "Index out of bounds", http.StatusBadRequest)
		return
	}

	people[idx] = Person{
		Name:    r.FormValue("name"),
		MetAt:   r.FormValue("met_at"),
		WorksAt: r.FormValue("works_at"),
		Fact:    r.FormValue("fact"),
	}

	err := templates.ExecuteTemplate(w, "person_row.html", struct {
		Index  int
		Person Person
	}{Index: idx, Person: people[idx]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	idx, _ := strconv.Atoi(r.FormValue("index"))
	if idx < 0 || idx >= len(people) {
		http.Error(w, "Index out of bounds", http.StatusBadRequest)
		return
	}

	people = append(people[:idx], people[idx+1:]...)
	w.WriteHeader(http.StatusOK)
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, fmt.Errorf("invalid dict call: odd number of args")
	}

	m := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, fmt.Errorf("dict keys must be strings")
		}
		m[key] = values[i+1]
	}
	return m, nil
}

func loadTemplates() *template.Template {
	funcMap := template.FuncMap{
		"dict": func(args ...interface{}) map[string]interface{} {
			m, _ := dict(args...)
			return m
		},
	}

	tmpl := template.New("").Funcs(funcMap)
	parsed, err := tmpl.ParseGlob(filepath.Join("templates", "*.html")) // or "cmd/server/templates/*.html"
	if err != nil {
		panic(err)
	}
	return parsed
}
