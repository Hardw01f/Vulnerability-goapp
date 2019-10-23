package logout

import (
	"net/http"
	"text/template"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		cookieSID := &http.Cookie{
			Name:  "SessionID",
			Value: "",
		}
		cookieUserName := &http.Cookie{
			Name:  "UserName",
			Value: "",
		}
		cookieUserID := &http.Cookie{
			Name:  "UserID",
			Value: "",
		}

		http.SetCookie(w, cookieSID)
		http.SetCookie(w, cookieUserName)
		http.SetCookie(w, cookieUserID)
		t, _ := template.ParseFiles("./views/public/logout.gtpl")
		t.Execute(w, nil)

	} else {
		http.NotFound(w, r)
	}
}
