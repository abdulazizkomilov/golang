package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"time"

	"web-cache/cache"

	"github.com/gorilla/mux"
)

const PORT = ":444"

type Page struct {
	Content string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}

	// Keshni yaratish
	cached := cache.NewCache("page", pageGUID)

	// Keshni tekshirish
	valid, cachedData := cached.Get()
	if valid {
		thisPage.Content = cachedData
		fmt.Fprintln(w, thisPage.Content)
		return
	}

	// Kesh mavjud bo'lmasa, ma'lumotni o'qish
	filePath := fmt.Sprintf("templates/%s.html", pageGUID)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	thisPage.Content = string(data)

	// Keshni yangilash
	cached.Set(thisPage.Content)

	t, err := template.New("page").ParseFiles("templates/blog.html")
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
		return
	}
	t.Execute(w, thisPage)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/page/{guid:[a-zA-Z0-9]+}", ServePage)

	server := &http.Server{
		Handler:      r,
		Addr:         PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server started at port", PORT)
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server failed:", err)
	}
}
