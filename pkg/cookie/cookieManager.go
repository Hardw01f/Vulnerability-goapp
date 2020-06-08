package cookie

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func CheckSessionsCount(uid int, sid string) bool {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("uid : ", uid)

	var count int

	if err := db.QueryRow("select count(sessionid) from vulnapp.sessions where uid=?", uid).Scan(&count); err != nil {
		fmt.Println(err)
	}

	fmt.Println("count is : ", count)

	if count != 0 {
		return false
	}
	return true

}

func ValidateCorrectCookie(uid int, sid string) bool {
	db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var sessionID string
	if err := db.QueryRow("select sessionid from vulnapp.sessions where uid=?", uid).Scan(&sessionID); err != nil {
		fmt.Println(err)
	}

	if sessionID == sid {
		return true
	} else {
		return false
	}

}

func CheckSessionID(r *http.Request) bool {
	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
		return false
	}

	userID, err := r.Cookie("UserID")
	if err != nil {
		fmt.Println(err)
		return false
	}

	if sessionID.Value == "" || userID.Value == "" {
		return false
	}

	uid, err := strconv.Atoi(userID.Value)
	if err != nil {
		fmt.Println(err)
	}

	if ValidateCorrectCookie(uid, sessionID.Value) {
		return true
	} else {
		return false
	}

}

func GetCookieValue(r *http.Request) (SessionID string, UserName string, UserID int, err error) {
	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
		return "", "", 0, err
	}

	userID, err := r.Cookie("UserID")
	if err != nil {
		fmt.Println(err)
	}
	uid, err := strconv.Atoi(userID.Value)
	if err != nil {
		fmt.Println(err)
		return "", "", 0, err
	}

	userName, err := r.Cookie("UserName")
	if err != nil {
		fmt.Println(err)
		return "", "", 0, err
	}

	return sessionID.Value, userName.Value, uid, nil
}

func CheckCookieOnlyLogin(r *http.Request) (userNameCookie string, sessionIDCookie string, err error) {
	userName, err := r.Cookie("UserName")
	if err != nil {
		fmt.Println(err)
	}

	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(userName, sessionID)

	if userName.Value == "" && sessionID.Value == "" {
		return "", "", errors.New("Cookie not exsit")
	} else {
		return userName.Value, sessionID.Value, nil
	}
}

func GetUserIDFromCookie(r *http.Request) (userNameCookie string, sessionIDCookie string, userIDCookie int, err error) {
	userName, err := r.Cookie("UserName")
	if err != nil {
		fmt.Println(err)
	}

	sessionID, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
	}

	if userName.Value == "" && sessionID.Value == "" {
		return "", "", 0, errors.New("not exist cookie")
	} else {
		decodeMail, err := base64.StdEncoding.DecodeString(sessionID.Value)
		if err != nil {
			fmt.Println(err)
		}
		mail := string(decodeMail)
		fmt.Println(mail)

		db, err := sql.Open("mysql", "root:rootwolf@tcp(mysql)/vulnapp")
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var userID int
		if err := db.QueryRow("select id from user where mail=?", mail).Scan(&userID); err != nil {
			fmt.Println("no set :", err)
		}

		log.Println(userID)
		return userName.Value, sessionID.Value, userID, nil
	}
}
