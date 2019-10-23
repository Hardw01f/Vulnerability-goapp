package register

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
)

type Person struct {
	UserName string
	Uid      int
}

func CheckUserDeplicate(mail string) bool {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var count int
	if err = db.QueryRow("select count(mail) from user where mail=?", mail).Scan(&count); err != nil {
		fmt.Println("db error : ", err)
	} else {
		fmt.Println(reflect.TypeOf(count))
		fmt.Println(count)
	}

	if count == 0 {
		return true
	}

	return false
}

func RegisterUser(r *http.Request) bool {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}

	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = db.Exec("insert into user (name,mail,age,passwd) value(?,?,?,?)", r.FormValue("name"), r.FormValue("mail"), age, r.FormValue("passwd"))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true

}

func NewUserRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./views/public/new.gtpl")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.FormValue("mail"))
		fmt.Println(r.FormValue("name"))
		fmt.Println(r.FormValue("age"))
		fmt.Println(r.FormValue("passwd"))

		if CheckUserDeplicate(r.FormValue("mail")) {

			if RegisterUser(r) {

				name := r.FormValue("name")
				mail := r.FormValue("mail")
				encodeMail := base64.StdEncoding.EncodeToString([]byte(mail))
				fmt.Println("register successful!!")
				cookieSID := &http.Cookie{
					Name:  "SessionID",
					Value: encodeMail,
				}
				cookieUserName := &http.Cookie{
					Name:  "UserName",
					Value: name,
				}

				http.SetCookie(w, cookieSID)
				http.SetCookie(w, cookieUserName)
				p := Person{UserName: name}
				t, _ := template.ParseFiles("./views/public/success_register.gtpl")
				t.Execute(w, p)
			}
		} else {
			t, _ := template.ParseFiles("./views/public/register_error.gtpl")
			t.Execute(w, nil)
		}

	} else {
		http.NotFound(w, r)
	}

}
