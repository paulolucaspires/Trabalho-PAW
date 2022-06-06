package main

import "net/http"

func (app *application) routes() *http.ServeMux {
  mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/projeto", app.perfil)
	mux.HandleFunc("/projeto/create", app.cadastro)
	fileServer := http.FileServer(http.Dir("./html"))
	mux.Handle("/html/", http.StripPrefix("/html", 
  fileServer))

  return mux
}