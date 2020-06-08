package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/hardw01f/Vulnerability-goapp/pkg/cookie"
)

type User struct {
	UserId   int
	UserName string
	Mail     string
	Age      int
	Image    string
	Address  string
	Animal   string
	Word     string
}

func ShowUserProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			_, _, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Println("CheckSessionID", err)
			}

			userName := GetUserName(uid)

			mail, age, err := GetUserInfos(uid)
			if err != nil {
				fmt.Println("GetUserInfos", err)
			}

			userImage, userAddress, userAnimal, userWord, err := GetOptUserDetails(uid)
			if err != nil {
				fmt.Println("GetOptUserDetails", err)
			}

			u := User{UserName: userName, Mail: mail, Age: age, Image: userImage, Address: userAddress, Animal: userAnimal, Word: userWord}
			t, _ := template.ParseFiles("./views/users/users.gtpl")
			t.Execute(w, u)
		} else {
			http.NotFound(w, nil)
		}

	} else {
		t, _ := template.ParseFiles("./views/error.gtpl")
		t.Execute(w, nil)
	}
}

func GetUserName(uid int) (Name string) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	defer db.Close()

	var name string

	if err := db.QueryRow("select name from vulnapp.user where id=?", uid).Scan(&name); err != nil {
		fmt.Println("%+v", err)
	}

	return name
}

func GetOptUserDetails(uid int) (usersimage string, usersaddress string, usersanimal string, usersword string, err error) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	defer db.Close()

	var userImage, address, animal, word string

	res, err := db.Query("select userImage,address,animal,word from vulnapp.userdetails where uid=?", uid)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	for res.Next() {
		if err := res.Scan(&userImage, &address, &animal, &word); err != nil {
			fmt.Printf("%+v", err)
		}
	}

	return userImage, address, animal, word, nil

}

func GetUserInfos(uid int) (userMail string, userAge int, err error) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Printf("%+v", err)
		return "", 0, err
	}
	defer db.Close()

	var mail string
	var age int

	res, err := db.Query("select mail,age from vulnapp.user where id=?", uid)
	if err != nil {
		fmt.Printf("%+v", err)
		return "", 0, err
	}

	for res.Next() {
		if err := res.Scan(&mail, &age); err != nil {
			fmt.Printf("%+v", err)
			return "", 0, err
		}
	}

	return mail, age, nil
}

func GetUserMoreDetails(uid int) (userImage string, Address string, Animal string, Word string, err error) {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Printf("%+v", err)
		return "", "", "", "", err
	}
	defer db.Close()

	var getUserImage, getUserAddress, getUserAnimal, getUserWord string

	res, err := db.Query("select userimage,address,animal,word from vulnapp.userdetails where uid=?", uid)
	if err != nil {
		fmt.Printf("%+v", err)
		return "", "", "", "", err
	}

	for res.Next() {
		if err := res.Scan(&getUserImage, &getUserAddress, &getUserAnimal, &getUserWord); err != nil {
			fmt.Printf("%+v", err)
			return "", "", "", "", err
		}
	}
	return getUserImage, getUserAddress, getUserAnimal, getUserWord, nil
}

func ShowUserModifyPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			_, userName, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Println("CheckSessionID", err)
			}
			mail, age, err := GetUserInfos(uid)
			if err != nil {
				fmt.Println("GetUserInfos", err)
			}
			userImage, userAddress, userAnimal, userWord, err := GetUserMoreDetails(uid)
			if err != nil {
				fmt.Println("GetUserMoreDetails", err)
			}

			u := User{UserName: userName, Mail: mail, Age: age, Image: userImage, Address: userAddress, Animal: userAnimal, Word: userWord}

			fmt.Println(u)
			t, _ := template.ParseFiles("./views/users/users_modify.gtpl")
			t.Execute(w, u)
		}

	} else {
		http.NotFound(w, nil)
	}
}

func ShowEditConfirm(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if cookie.CheckSessionID(r) {
			fmt.Println(r.Referer())
			userName := r.FormValue("username")
			age := r.FormValue("age")
			mail := r.FormValue("mail")
			fmt.Println("mail : ", mail)
			address := r.FormValue("address")
			animal := r.FormValue("animal")
			word := r.FormValue("word")

			userAge, err := strconv.Atoi(age)
			if err != nil {
				fmt.Printf("%+v\n", err)
			}

			_, _, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Printf("%+v\n", err)
				http.NotFound(w, nil)
				return
			}

			userMail, _, err := GetUserInfos(uid)
			if err != nil {
				fmt.Printf("%+v", err)
			}

			userImage, _, _, _, err := GetUserMoreDetails(uid)
			if err != nil {
				fmt.Printf("%+v", err)
				http.NotFound(w, nil)
				return
			}

			u := User{UserName: userName, Mail: userMail, Age: userAge, Image: userImage, Address: address, Animal: animal, Word: word}

			t, _ := template.ParseFiles("./views/users/users_confirm.gtpl")
			t.Execute(w, u)
		}
	} else {
		http.NotFound(w, nil)
	}
}

func CheckRecodeUserDetails(uid int) bool {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	defer db.Close()

	var count int
	if err = db.QueryRow("select uid from userdetails where uid = ?", uid).Scan(&count); err != nil {
		fmt.Println("%+v\n", err)
		return false
	}

	if count == 0 {
		return false
	} else {
		return true
	}

}

func UpdateUserDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if cookie.CheckSessionID(r) {
			userName := r.FormValue("username")
			age := r.FormValue("age")
			address := r.FormValue("address")
			animal := r.FormValue("animal")
			word := r.FormValue("word")

			db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
			if err != nil {
				fmt.Printf("%+v", err)
			}
			defer db.Close()

			_, _, uid, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Printf("%+v\n", err)
				http.NotFound(w, nil)
				return
			}

			_, err = db.Exec("update vulnapp.user set name = ?, age = ? where id = ?", userName, age, uid)
			if err != nil {
				fmt.Printf("%+v\n", err)
				http.NotFound(w, nil)
				return
			}

			if CheckRecodeUserDetails(uid) {

				_, err = db.Exec("update vulnapp.userdetails set address = ?, animal = ?, word = ? where uid = ?", address, animal, word, uid)
				if err != nil {
					fmt.Printf("%+v\n", err)
					http.NotFound(w, nil)
					return
				}

			} else {
				_, err = db.Exec("insert into vulnapp.userdetails (uid,userimage,address,animal,word) values (?,?,?,?,?)", uid, "noimage.png", address, animal, word)
				if err != nil {
					fmt.Printf("%+v\n", err)
					http.NotFound(w, nil)
					return
				}

			}

			http.Redirect(w, r, "/profile", 301)

		}

	} else {
		http.NotFound(w, nil)
	}
}

func PasswdChange(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		if cookie.CheckSessionID(r) {
			t, _ := template.ParseFiles("./views/passwd/passwd_modify.gtpl")
			t.Execute(w, nil)

		}
	} else if r.Method == "POST" {
		if cookie.CheckSessionID(r) {
			_, _, _, err := cookie.GetCookieValue(r)
			if err != nil {
				fmt.Printf("%+v\n", err)
			}

			newPasswd := r.FormValue("newpasswd")
			confirmPassword := r.FormValue("confirmpasswd")

			if newPasswd == confirmPassword {

				t, _ := template.ParseFiles("./views/passwd/hiddenchangepasswd.gtpl")
				h := map[string]string{"Passwd": newPasswd, "Confirm": confirmPassword}

				t.Execute(w, h)

			} else {
				http.NotFound(w, nil)
			}

		}

	} else {
		http.NotFound(w, nil)
	}
}

func ConfirmPasswdChange(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		if cookie.CheckSessionID(r) {
			if r.Referer() == "http://localhost:9090/profile/changepasswd" {
				_, _, uid, err := cookie.GetCookieValue(r)
				if err != nil {
					fmt.Printf("%+v\n", err)
				}

				newPasswd := r.FormValue("passwd")
				confirmPassword := r.FormValue("confirm")

				fmt.Println(newPasswd, confirmPassword, uid)

				db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
				if err != nil {
					fmt.Printf("%+v", err)
				}
				defer db.Close()

				_, err = db.Exec("update vulnapp.user set passwd = ? where id = ?", newPasswd, uid)
				if err != nil {
					fmt.Printf("%+v\n", err)
					http.NotFound(w, nil)
					return
				}

				t, _ := template.ParseFiles("./views/passwd/passwdchengedcomplete.gtpl")
				t.Execute(w, nil)

			} else {
				http.NotFound(w, nil)
			}

		}

	} else {
		http.NotFound(w, nil)
	}
}
