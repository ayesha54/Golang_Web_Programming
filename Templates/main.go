package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	fname string
	lname string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template_files/*"))

}
func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "index.html", "John")
	if err != nil {
		log.Fatalln(err)
	}

	//fruits := []string{"apple", "banana", "mango"}
	//ages := map[string]int{"john": 45, "jonas": 42}
	p1 := person{
		"jane",
		"hilton",
	}

	err = tpl.ExecuteTemplate(os.Stdout, "home.html", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
