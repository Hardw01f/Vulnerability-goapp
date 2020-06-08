package login

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/hardw01f/Vulnerability-goapp/pkg/cookie"
)

type Person struct {
	UserName string
}

func outErrorPage(w http.ResponseWriter) {
	t, _ := template.ParseFiles("./views/public/error.gtpl")
	t.Execute(w, nil)
}

func SearchID(mail string) int {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "select id from user where mail=?"
	res, err := db.Query(sql, mail)
	if err != nil {
		log.Fatal(err)
	}

	var id int

	for res.Next() {
		err := res.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("ID :", id)
	}

	fmt.Println(id)
	return id
}

func CheckPasswd(id int, passwd string) string {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	sql := "select name from user where id=? and passwd=?"
	res, err := db.Query(sql, id, passwd)
	if err != nil {
		log.Fatal(err)
	}

	for res.Next() {
		err := res.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}

	return name
}

func isZeroString(formstr string) bool {
	//fmt.Println("len: ", len(formstr))
	if len(formstr) == 0 {
		return false
	}
	return true
}

func StoreSID(uid int, sid string) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if cookie.CheckSessionsCount(uid, sid) {
		_, err = db.Exec("insert into sessions(uid,sessionid) value (?,?)", uid, sid)
		if err != nil {
			fmt.Println(err)
		}
	} else {
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method ", r.Method)

	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			http.Redirect(w, r, "/top", 302)
		} else {
			t, _ := template.ParseFiles("./views/public/login.gtpl")
			t.Execute(w, nil)
		}
	} else if r.Method == "POST" {

		r.ParseForm()
		if isZeroString(r.FormValue("mail")) && isZeroString(r.FormValue("passwd")) {
			fmt.Println("passwd", r.Form["passwd"])
			fmt.Println("mail", r.Form["mail"])

			mail := r.FormValue("mail")
			id := SearchID(mail)

			if id != 0 {
				passwd := r.FormValue("passwd")
				name := CheckPasswd(id, passwd)

				if name != "" {
					fmt.Println(name)
					t, _ := template.ParseFiles("./views/public/logined.gtpl")
					encodeMail := base64.StdEncoding.EncodeToString([]byte(mail))
					fmt.Println(encodeMail)
					cookieSID := &http.Cookie{
						Name:  "SessionID",
						Value: encodeMail,
					}
					cookieUserName := &http.Cookie{
						Name:  "UserName",
						Value: name,
					}
					StoreSID(id, encodeMail)
					http.SetCookie(w, cookieUserName)
					http.SetCookie(w, cookieSID)
					p := Person{UserName: name}
					t.Execute(w, p)
				} else {
					fmt.Println(name)
					t, _ := template.ParseFiles("./views/public/error.gtpl")
					t.Execute(w, nil)
				}
			} else {
				t, _ := template.ParseFiles("./views/public/error.gtpl")
				t.Execute(w, nil)
			}

		} else {
			fmt.Println("username or passwd are empty")
			outErrorPage(w)
		}
	} else {
		http.NotFound(w, nil)
	}
}
