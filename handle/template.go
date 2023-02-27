package handle

import (
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

type Arg struct {
	Boolean    bool
	DaysOfWeek []string
}

func Template(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("wwwroot/template.html")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("aaa")
		return
	}
	rng := rand.New(rand.NewSource(time.Now().Unix()))
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	a := Arg{
		Boolean:    rng.Intn(10) > 5,
		DaysOfWeek: daysOfWeek,
	}
	t.Execute(w, a)
}
