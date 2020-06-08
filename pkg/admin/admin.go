package admin

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/xerrors"

	"github.com/hardw01f/Vulnerability-goapp/pkg/cookie"
)

type Lists struct {
	Uid       string
	UserName  string
	UserLists []string
}

func GetRandString() string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func StoreAdminSID(adminSessionID string) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("insert into adminsessions(adminsessionid) values(?)", adminSessionID)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

}

func GetAdminSid(adminSessionCookie string) (results string, err error) {
	commandLine := "mysql -h mysql -u root -prootwolf -e 'select adminsid from vulnapp.adminsessions where adminsessionid=\"" + adminSessionCookie + "\";'"

	res, err := exec.Command("sh", "-c", commandLine).Output()
	if err != nil {
		fmt.Println(err)
	}

	results = string(res)

	if results != "" {
		return results, nil
	}

	err = xerrors.New("recode was not set")

	return "", err
}

func ShowAdminLogIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			t, _ := template.ParseFiles("./views/admin/adminlogin.gtpl")
			t.Execute(w, nil)
		} else {
			http.Redirect(w, r, "/login", 302)
		}
	} else {
		http.NotFound(w, nil)
	}
}

func Confirm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		requestMail := r.FormValue("adminmail")
		requestPasswd := r.FormValue("adminpasswd")

		fmt.Println(requestMail, ":", requestPasswd)

		cmd := "mysql -h mysql -u root -prootwolf -e 'select adminid from vulnapp.admins where mail=\"" + requestMail + "\" and passwd=\"" + requestPasswd + "\";'"

		fmt.Println(cmd)

		res, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println("err : ", err)
		}

		if string(res) == "" {
			fmt.Println("not")
			t, _ := template.ParseFiles("./views/admin/failedauthentication.gtpl")
			t.Execute(w, nil)
			return
		}

		fmt.Println(string(res))
		fmt.Println("success")

		adminSessionID := GetRandString()
		fmt.Println(adminSessionID)

		adminSID := &http.Cookie{
			Name:  "adminSID",
			Value: adminSessionID,
		}
		http.SetCookie(w, adminSID)

		StoreAdminSID(adminSessionID)

		t, _ := template.ParseFiles("./views/admin/successauthentication.gtpl")
		t.Execute(w, nil)

	} else {
		http.NotFound(w, nil)
	}
}

func ShowAdminPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		adminSID, err := r.Cookie("adminSID")
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
		fmt.Println(adminSID.Value)

		adminUid, err := GetAdminSid(adminSID.Value)
		if err != nil {
			fmt.Println("not authentication")
			t, _ := template.ParseFiles("./views/admin/failedauthentication.gtpl")
			t.Execute(w, nil)
			return
		}

		fmt.Println(adminUid)

		uid, err := r.Cookie("UserID")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(uid)

		userName, err := r.Cookie("UserName")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(userName.Value)

		cmd := "mysql -h mysql -u root -prootwolf -e 'select id,name,mail,age,created_at,updated_at from vulnapp.user where name not in (\"" + userName.Value + "\");'"

		fmt.Println(cmd)

		res, err := exec.Command("sh", "-c", cmd).Output()
		if err != nil {
			fmt.Println("err : ", err)
		}

		splitedRes := strings.Split(string(res), "\n")
		fmt.Println(splitedRes)

		p := Lists{Uid: uid.Value, UserName: userName.Value, UserLists: splitedRes}

		fmt.Println(p)

		t, _ := template.ParseFiles("./views/admin/userlists.gtpl")
		t.Execute(w, p)

	} else {
		http.NotFound(w, nil)
	}
}
