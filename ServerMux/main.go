package main

import "net/http"

func main() {
	mux := http.NewServeMux()//criar um server mux
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {//personalizando a rota
	// 	w.Write([]byte("Hello World"))
	// })
	mux.HandleFunc("/", HomeHandler)//definir a rota
	mux.Handle("/blog", blog{title: "My blog"})
	http.ListenAndServe(":8080", mux)//definir a porta
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}