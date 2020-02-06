package main

import (
	"html/template"
	"log"
    "fmt"
	"net/http"
)

func CSRFTrap(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
	t, _ := template.ParseFiles("./Trap.gtpl")
	t.Execute(w, r)
    } else {
        http.NotFound(w,nil)
    }
}

func PasswdCSRF(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("./PasswdCSRF.gtpl")
    t.Execute(w,nil)
}

func DetailCSRF(w http.ResponseWriter, r *http.Request){
    t, _ := template.ParseFiles("./DetailCSRF.gtpl")
    t.Execute(w,nil)
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/csrftrap", CSRFTrap)
	http.HandleFunc("/detailCSRF", DetailCSRF)
	http.HandleFunc("/passwdCSRF", PasswdCSRF)
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
    fmt.Println("ListenPost : 3030")
    }
