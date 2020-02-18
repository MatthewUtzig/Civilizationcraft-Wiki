package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"html/template"
)

const debug_mode bool = true;

var tmpl *template.Template


//type TowerStrcture struct {
//	Structure
//	Tower
//}

func main()  {
	route := mux.NewRouter()

	//handlers for css and images
	cssHandler := http.FileServer(http.Dir("./css/"))
	imagesHandler := http.FileServer(http.Dir("./images/"))
	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/images/", http.StripPrefix("/images/", imagesHandler))

	route.HandleFunc("/", HomePage)
	route.HandleFunc("/culture", CulturePage)

	http.Handle("/", route)

	fmt.Println("Starting server")
	log.Fatalln(http.ListenAndServe(":80", nil))
}


func HomePage(w http.ResponseWriter, r *http.Request) {
	LoadTemplates()
	tmpl.ExecuteTemplate(w,"main.gohtml", nil)
}

func CulturePage(w http.ResponseWriter, r *http.Request) {
	LoadTemplates()
	tmpl.ExecuteTemplate(w,"culture.gohtml", nil)
}


/*
* Reloads templates if debug mode is on
 */
func ReloadTemplates() {
	if (debug_mode) {
		LoadTemplates()
	}
}

func LoadTemplates() {
	var err error
	tmpl, err = template.New("junk").ParseGlob("templates/*")

	if (err != nil) {
		log.Fatalln("Error parsing templates", err)
	}
}
