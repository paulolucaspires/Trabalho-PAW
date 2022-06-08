package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(rw)
		return
	}

  formulario, err := app.formulario.Latest()
  if err != nil{
    app.serverError(rw, err)
    return
  }

	files := []string{
		"./html/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, formulario)
	if err != nil {
		app.serverError(rw, err)
		return
	}

}

func (app *application) perfil(rw http.ResponseWriter, r *http.Request) {
	name, err := strconv.Atoi(r.URL.Query().Get("name"))
	if err != nil  {
		app.notFound(rw)
		return
	}

  s, err := app.formualrio.Get(name)
  if err == models.ErrNoRecord {
    app.notFound(rw)
    return
  }else if err != nil{
    app.serverError(rw, err)
    return
  }
  
}

func (app *application) cadastro(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		rw.Header().Set("Allow", "POST")
		app.clientError(rw, http.StatusMethodNotAllowed)
		return
	}

	title := "Aula de hoje"
	content := "Tentando lidar com o banco de dados"
	expire := "7"

	id, err := app.formulario.Insert(title, content, expire)
	if err != nil {
		app.serverError(rw, err)
		return
	}
  files := []string{
		"./html/home.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(rw, err)
		return
	}
	err = ts.Execute(rw, formulario)
	if err != nil {
		app.serverError(rw, err)
		return
	}

	http.Redirect(rw, r, fmt.Sprintf("/formulario?id=%d", id), http.StatusSeeOther)
}