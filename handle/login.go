package handle

import (
	"fmt"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	// login 处理细节
	t, err := template.ParseFiles("wwwroot/login.html")
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	t.Execute(w, fmt.Sprintf("login successfully. Welcome %v", r.PostFormValue("username")))
}
