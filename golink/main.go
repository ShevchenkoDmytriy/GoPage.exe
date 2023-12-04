package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// type Users struct {
// 	Name       string
// 	Age        uint
// 	Money      int
// 	Avr_grades float64
// 	Happiness  float64
// 	Hobbies    []string
// }

// func (u Users) getAllInfo() string {
// 	return fmt.Sprintf("%s %d %d %f %f", u.Name, u.Age, u.Money, u.Avr_grades, u.Happiness)
// }
// func (u *Users) setNewName(name string) {
// 	u.Name = name
// }
// func home_page(w http.ResponseWriter, r *http.Request) {
// 	bob := Users{"Bob", 25, -50, 4.2, 0.8, []string{"q", "w", "e", "r"}}
// 	// fmt.Fprintf(w, "User:%s %d %d %f %f\n", bob.name, bob.age, bob.money, bob.avr_grades, bob.happiness)
// 	// bob.setNewName("Tom")
// 	// fmt.Fprintf(w, `<h1>Main</h1>`)
// 	tmpl, _ := template.ParseFiles("template/homepage.html")
// 	tmpl.Execute(w, bob)
// }

// func contacts_page(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Second page")
// }
// func handleRequest() {
// 	http.HandleFunc("/", home_page)
// 	http.HandleFunc("/contacts/", contacts_page)

// 	http.ListenAndServe(":8080", nil)
// }
// func main() {

//		handleRequest()
//	}
type User struct {
	Name string `json:"Name"`
	Age  uint   `json:"Age"`
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err.Error)
	}
	defer db.Close()
	// insert, err := db.Query("INSERT INTO `users`(`Name`,`Age`) VALUES('Tom',20)")
	// if err != nil {
	// 	panic(err.Error)
	// }
	// defer insert.Close()

	send, err := db.Query("SELECT `Name`,`Age` FROM `users`")
	if err != nil {
		panic(err.Error)
	}
	for send.Next() {
		var users User
		err = send.Scan(&users.Name, &users.Age)
		if err != nil {
			panic(err.Error)
		}

		fmt.Println(fmt.Sprintf("Users: %s \tAge: %d", users.Name, users.Age))
	}

}
