package main

import (
	"log"
	"os"
	"text/template"
	"time"

	"gopl.io/ch4/github"
)


const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | dayAgo}} days
{{end}}`

func daysAgo(t time.Time)int{
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("issuelist").
  Funcs(template.FuncMap{"daysAgo":daysAgo}).
  Parse(templ))

  func main(){

  }

func noMust(){
	report, err := template.New("report")
	
}










