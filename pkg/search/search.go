package search

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"text/template"
)

type Results struct {
	StringResults []string
}

func SearchPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		searchWord := r.FormValue("post")
		fmt.Println("value : ", searchWord)

		execSql := "mysql -h mysql -u root -prootwolf -e 'select post,created_at from vulnapp.posts where post like \"%" + searchWord + "%\"'"

		fmt.Println(execSql)

		sqlRes, err := exec.Command("sh", "-c", execSql).Output()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(sqlRes))

		splitedRes := strings.Split(string(sqlRes), "\n")

		var results []string

		for i, v := range splitedRes {
			if i > 0 {
				fmt.Println(i, v)
				results = append(results, v)
			}
		}

		fmt.Println(len(results))

		var removedExtraSpace []string
		for i := 0; i < len(results)-1; i++ {
			removedExtraSpace = append(removedExtraSpace, results[i])
		}
		fmt.Println(len(removedExtraSpace))

		p := Results{StringResults: removedExtraSpace}

		fmt.Println(p)

		t, _ := template.ParseFiles("./views/search/searchpost.gtpl")
		t.Execute(w, p)

	} else {
		http.NotFound(w, nil)
	}
}
