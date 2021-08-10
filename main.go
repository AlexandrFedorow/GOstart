package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type User struct {
	Name string
	Age uint16
	Money int16
	Avg_grades float64
	Happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User name is: %s \n He is: %d \n" +
		"And he has money: %d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request)  {
	bob := User{"Bob", 15, 50, 1.5, 5}
	//bob.setNewName("jija")
	//fmt.Fprintf(w, bob.getAllInfo())
	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "contct")
}

func handleRequest()  {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8000", nil)
}

func main()  {
	handleRequest()
}
