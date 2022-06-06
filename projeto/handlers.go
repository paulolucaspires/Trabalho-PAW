package main

import (
	"fmt"
  "html/template"
	"log"
	"net/http"
	"strconv"
)

func (app *application)home(rw http.ResponseWriter, r *http.Request){
   if r.URL.Path != "/"{
    app.NotFound(rw, r)
    return
  } 
  rw.Write([]byte("Formulário de Contato"))
}


func (app *application)perfil(rw http.ResponseWriter, r *http.Request){
  id,err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.NotFound(rw, r)
    return
  }
  fmt.Fprintf(rw, "Exibir o perfil de ID: %d", id)
}

func (app *application)cadastro(rw http.ResponseWriter, r *http.Request){
  if r.Method != "POST"{
    rw.Header().Set("Allow","POST")
    app.clientError(rw, http.StatusMethodNotAllowed)
    return
  }

  rw.Write([]byte("Criar novo usuario"))
}