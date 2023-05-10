package main

import (
    "html/template"
    "net/http"
)

type Item struct {
    ID    int
    Name  string
    Price float64
}

// Database to store items
var items = map[int]*Item{}

// Counter for generating IDs
var counter = 1

// Handler function for displaying the form to create a new item
func newItemFormHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("new-item.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := t.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Handler function for creating a new item
func newItemHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()

    name := r.FormValue("name")
    price := r.FormValue("price")

    item := &Item{
        ID:    counter,
        Name:  name,
        Price: parseFloat(price),
    }

    items[counter] = item
    counter++

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Handler function for displaying the list of items
func listItemsHandler(w http.ResponseWriter, r *http.Request) {
    t, err := template.ParseFiles("list-items.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := t.Execute(w, items); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func main() {
    // Set up routes
    http.HandleFunc("/", listItemsHandler)
    http.HandleFunc("/new-item", newItemFormHandler)
    http.HandleFunc("/create-item", newItemHandler)

    // Start server
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
