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

func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	//Permet de recuperer la requete en frond pour la mettre dasn la db

	// var register User
	// decoder := json.NewDecoder(r.Body)
	// decoder.Decode(&register)
	// fmt.Println(register)

	// insert := `INSERT INTO bobato.session (USR_ID, TOKEN) VALUES (` + strconv.Itoa(session.UserID) + `,"` + session.Token + `");`

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
	// http.HandleFunc("/api/login", test)
	http.HandleFunc("/api/pierre", pierresHandler)
	http.HandleFunc("/api/pierre/", pierreHandler)
	// http.HandleFunc("/api/cart", test)
	http.HandleFunc("/api/user", userHandler)

	log.Fatal(http.ListenAndServe(":55", nil))

}
