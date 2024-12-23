package main

import (
	"html/template"
	"net/http"
)

var posts = map[string]string{
	"post1": "This is the first blog post.",
	"post2": "Another interesting post!",
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("post.html")
	postID := r.URL.Path[len("/post/"):]
	content, ok := posts[postID]
	if ok {
		tmpl.Execute(w, content)
	} else {
		http.NotFound(w, r)
	}
}

// func main() {
// 	http.HandleFunc("/post/", postHandler)
// 	fmt.Println("Starting server on :8080")
// 	http.ListenAndServe(":8080", nil)
// }
