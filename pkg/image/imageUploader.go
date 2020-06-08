package uploader

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/hardw01f/Vulnerability-goapp/pkg/cookie"
	"github.com/hardw01f/Vulnerability-goapp/pkg/user"
)

type User struct {
	Image string
}

func ShowImageChangePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			_, _, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Printf("%+v\n", err)
				http.NotFound(w, nil)
				return
			}
			userImage, _, _, _, err := user.GetOptUserDetails(uid)
			if err != nil {
				fmt.Printf("%+v\n", err)
				http.NotFound(w, nil)
				return
			}

			u := User{Image: userImage}
			t, _ := template.ParseFiles("./views/users/users_image.gtpl")
			t.Execute(w, u)
		}

	} else {
		http.NotFound(w, nil)
	}
}

func UpdateDatabase(r *http.Request, imageName string) error {
	_, _, uid, err := cookie.GetCookieValue(r)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}

	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}
	defer db.Close()

	_, err = db.Exec("update vulnapp.userdetails set userimage = ? where uid = ?", imageName, uid)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return err
	}

	return nil

}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(32 << 20) // maxMemory
		if err != nil {
			fmt.Printf("%+v\n", err)
			return
		}
		if cookie.CheckSessionID(r) {
			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				fmt.Printf("%+v\n", err)
				return
			}
			defer file.Close()
			f, err := os.OpenFile("./assets/img/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Printf("%+v\n", err)
				return
			}
			defer f.Close()
			io.Copy(f, file)
			UpdateDatabase(r, handler.Filename)

			http.Redirect(w, r, "/profile", 301)
		}
	} else {
		http.NotFound(w, nil)
	}

}
