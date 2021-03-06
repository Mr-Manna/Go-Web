package main

import (
	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

type Page struct{
	Name string
	DBStatus bool
}

func main(){
	templates := template.Must(template.ParseFiles("templates/index.html"))
	fmt.Println("Server Is On")
	//  connecting Database
	db, _ := sql.Open("sqlite3","dev.db")

	http.HandleFunc("/", func( w http.ResponseWriter, r *http.Request ){
		p := Page{ Name: "Gopher"}

		if name := r.FormValue("name"); name != " " {
			p.Name = name
			}
		// http://localhost:8000/?name=name to pass the variable
			p.DBStatus = db.Ping() == nil

		if err := templates.ExecuteTemplate( w, "index.html", p); err!= nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// fmt.Fprintf( w, " Hello World! ")
	})
	
	fmt.Println(http.ListenAndServe( ":8888",nil))
	}