package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db, _ = sql.Open("mysql", "root:root@tcp(localhost)/labaiepierre")

type User struct {
	ID        int    `json:"ID"`
	Name      string `json:"name"`
	Firstname string `json:"firstname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	PP        string `json:"PP"`
	Cart_ID   int    `json:"cart_ID"`
	Birthday  string `json:"birthday"`
}

type Session struct {
	UserID   int
	UserName string
	Token    string
}

type Pierre struct {
	ID                 int    `json:"ID"`
	Pierre_name        string `json:"pierre_name"`
	Pierre_description string `json:"pierre_description"`
	Pierre_price       int    `json:"pierre_price"`
	Categorie          string `json:"categorie"`
	Avis               Avis
}
type Avis struct {
	ID        int    `json:"ID"`
	Pierre_ID int    `json:"pierre_ID"`
	User_ID   int    `json:"user_ID"`
	Note      string `json:"note"`
	Text      string `json:"text"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
}

func getPierres() []Pierre {
	var query, _ = db.Query("SELECT * FROM labaiepierre.pierre")
	var pierres []Pierre
	for query.Next() {
		var pierre Pierre
		query.Scan(&pierre.ID, &pierre.Pierre_name, &pierre.Pierre_description, &pierre.Pierre_price, &pierre.Categorie)
		pierres = append(pierres, pierre)
	}
	return pierres
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	//Permet de recuperer la requete en frond pour la mettre dasn la db

	var register User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&register)
	fmt.Println(register)
	insert := `INSERT INTO labaiepierre.user (NAME, FIRSTNAME, PASSWORD, EMAIL, BIRTHDAY) VALUES ("` + register.Name + `","` + register.Firstname + `","` + register.Password + `","` + register.Email + `","` + register.Birthday + `")`
	db.Query(insert)
	selectID := `SELECT ID FROM labaiepierre.user WHERE EMAIL="` + register.Email + `"`
	IDUsr := db.QueryRow(selectID)
	IDUsr.Scan(&register.ID)
	fmt.Println(register.ID)
	insertCart := `INSERT INTO labaiepierre.cart (USER_ID) VALUES ("` + strconv.Itoa(register.ID) + `")`
	db.Query(insertCart)
	selectCartID := `SELECT ID FROM labaiepierre.cart WHERE USER_ID="` + strconv.Itoa(register.ID) + `"`
	IDCart := db.QueryRow(selectCartID)
	IDCart.Scan(&register.Cart_ID)
	fmt.Println(register.Cart_ID)
	insertCartID := `UPDATE labaiepierre.user SET CART_ID="` + strconv.Itoa(register.Cart_ID) + `" WHERE ID="` + strconv.Itoa(register.ID) + `"`
	db.Query(insertCartID)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var user User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	emailVar := `SELECT ID, NAME, FIRSTNAME, PASSWORD, EMAIL, CART_ID, BIRTHDAY FROM labaiepierre.user WHERE EMAIL="` + user.Email + `" AND PASSWORD="` + user.Password + `"`
	fmt.Println(emailVar)
	var getRaw = db.QueryRow(emailVar)
	getRaw.Scan(&user.ID, &user.Name, &user.Firstname, &user.Password, &user.Email, &user.Cart_ID, &user.Birthday)
	fmt.Println(user)
}

// func cartHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Content-Type", "application/json")
// 	var pierre Pierre
// 	decoder := json.NewDecoder(r.Body)
// 	decoder.Decode(&pierre)
// 	pierreVar := `INSERT INTO labaiepierre.cart  (USER_ID, PIERRE_ID) VALUES ("` + strconv.Itoa(pierre.ID) + `")`
// }

func pierresHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	a, _ := json.Marshal(getPierres())
	w.Write(a)
}

func pierreHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	pathID := r.URL.Path
	pathID = path.Base(pathID)
	pathIDint, _ := strconv.Atoi(pathID)
	getPierresVar := getPierres()
	a, _ := json.Marshal(getPierresVar[pathIDint-1])
	w.Write(a)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var query = db.QueryRow("SELECT * FROM labaiepierre.user WHERE ID=2")
	var user User
	query.Scan(&user.ID, &user.Name, &user.Firstname, &user.Password, &user.Email, &user.PP, user.Cart_ID, user.Birthday)
	fmt.Println("ta grosse darone", user)
	a, _ := json.Marshal(user)
	w.Write(a)

}

func main() {
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/api/register", registerHandler)
	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/pierre", pierresHandler)
	http.HandleFunc("/api/pierre/", pierreHandler)
	http.HandleFunc("/api/cart", cartHandler)
	http.HandleFunc("/api/user", userHandler)

	log.Fatal(http.ListenAndServe(":55", nil))

}
