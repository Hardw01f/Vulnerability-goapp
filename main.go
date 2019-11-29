package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	//login "github.com/Snow-HardWolf/Vulnerability-goapp/pkg/login"

	"./pkg/admin"
	"./pkg/cookie"
	"./pkg/image"
	"./pkg/login"
	"./pkg/logout"
	"./pkg/post"
	"./pkg/register"
	"./pkg/search"
	"./pkg/user"
)

type Person struct {
	UserName string
}

func sayYourName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println("r.Form", r.Form)
	fmt.Println("r.Form[name]", r.Form["name"])
	var Name string
	for k, v := range r.Form {
		fmt.Println("key:", k)
		Name = strings.Join(v, ",")
	}
	fmt.Println(Name)
	fmt.Fprintf(w, Name)
}

func showUserTopPage(w http.ResponseWriter, r *http.Request) {
	userName, sessionID, userID, err := cookie.GetUserIDFromCookie(r)
	if err != nil {
		fmt.Println(err)
	}

	if cookie.CheckSessionsCount(userID, sessionID) {
		login.StoreSID(userID, sessionID)
	} else {
		fmt.Println("not register sessionID")
	}

	if sessionID == "" {
		fmt.Println("sid not exist")
		t, _ := template.ParseFiles("./views/public/error.gtpl")
		t.Execute(w, nil)
	} else {

		if r.Method == "GET" {

			if userID != 0 {
				uid := strconv.Itoa(userID)
				cookieUserID := &http.Cookie{
					Name:  "UserID",
					Value: uid,
				}

				deleAdminID := &http.Cookie{
					Name:  "adminSID",
					Value: "",
				}

				http.SetCookie(w, cookieUserID)
				http.SetCookie(w, deleAdminID)
				p := Person{UserName: userName}
				t, _ := template.ParseFiles("./views/public/top.gtpl")
				t.Execute(w, p)
			}

		} else {
			http.NotFound(w, r)
		}
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	userName, sid, err := cookie.CheckCookieOnlyLogin(r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(userName, sid)
	_, _, userID, err := cookie.GetUserIDFromCookie(r)
	if err != nil {
		fmt.Println(err)
	}
	if cookie.CheckSessionsCount(userID, sid) {
		fmt.Println("session count true")
	} else {
		fmt.Println("session count false")
	}

	if cookie.CheckSessionID(r) {
		fmt.Println("checkCookie true")
	} else {
		fmt.Println("checkCookie false")
	}

	t, _ := template.ParseFiles("./views/test.gtpl")
	t.Execute(w, nil)

}

func DBDetails(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./views/public/DBDetails.gtpl")
	t.Execute(w, nil)
}

func Hints(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./views/hints/hints.gtpl")
	t.Execute(w, nil)
}

func main() {
	var portNum = flag.String("p", "80", "Specify application server listening port")
	flag.Parse()
	fmt.Println("Vulnapp server listening : " + *portNum)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	http.HandleFunc("/", sayYourName)
	http.HandleFunc("/test", test)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/logout", logout.Logout)
	http.HandleFunc("/new", register.NewUserRegister)
	http.HandleFunc("/top", showUserTopPage)
	http.HandleFunc("/profile", user.ShowUserProfile)
	http.HandleFunc("/profile/edit", user.ShowUserModifyPage)
	http.HandleFunc("/profile/edit/confirm", user.ShowEditConfirm)
	http.HandleFunc("/profile/edit/update", user.UpdateUserDetails)
	http.HandleFunc("/profile/changepasswd", user.PasswdChange)
	http.HandleFunc("/profile/compchangepasswd", user.ConfirmPasswdChange)
	http.HandleFunc("/profile/edit/image", uploader.ShowImageChangePage)
	http.HandleFunc("/profile/edit/upload", uploader.UploadImage)
	http.HandleFunc("/post", post.ShowAddPostPage)
	http.HandleFunc("/timeline", post.ShowTimeline)
	http.HandleFunc("/timeline/searchpost", search.SearchPosts)
	http.HandleFunc("/hints", Hints)
	http.HandleFunc("/db", DBDetails)
	http.HandleFunc("/adminlogin", admin.ShowAdminLogIn)
	http.HandleFunc("/adminconfirm", admin.Confirm)
	http.HandleFunc("/adminusers", admin.ShowAdminPage)
	err := http.ListenAndServe(":"+*portNum, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
