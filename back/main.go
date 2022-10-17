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
	Birthday  string `json:"birthday"`
}
type Session struct {
	UserID        int    `json:"user_ID"`
	UserName      string `json:"user_name"`
	UserFirstname string `json:"user_firstname"`
	Token         string
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

type Cart struct {
	ID           int `json:"ID"`
	User_ID      int `json:"user_ID"`
	Pierre_ID    int `json:"pierre_ID"`
	Pierre_price int `json:"pierre_price"`
}

//

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
	insert := `INSERT INTO labaiepierre.user (NAME, FIRSTNAME, PASSWORD, EMAIL, BIRTHDAY) VALUES ("` + register.Name + `","` + register.Firstname + `","` + register.Password + `","` + register.Email + `","` + register.Birthday + `")`
	db.Query(insert)
	selectID := `SELECT ID FROM labaiepierre.user WHERE EMAIL="` + register.Email + `"`
	fmt.Println(selectID)
	IDUsr := db.QueryRow(selectID)
	IDUsr.Scan(&register.ID)
	fmt.Println(register)
	a, _ := json.Marshal(register)
	w.Write(a)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	var user User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&user)
	emailVar := `SELECT ID, NAME, FIRSTNAME, PASSWORD, EMAIL, BIRTHDAY FROM labaiepierre.user WHERE EMAIL="` + user.Email + `" AND PASSWORD="` + user.Password + `"`
	fmt.Println(emailVar)
	var getRaw = db.QueryRow(emailVar)
	getRaw.Scan(&user.ID, &user.Name, &user.Firstname, &user.Password, &user.Email, &user.Birthday)
	fmt.Println(user)
	a, _ := json.Marshal(user)
	w.Write(a)

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

	var usr User
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&usr)
	fmt.Println(usr)

	var query = db.QueryRow("SELECT * FROM labaiepierre.user WHERE ID=" + strconv.Itoa(usr.ID))
	var user User
	query.Scan(&user.ID, &user.Name, &user.Firstname, &user.Password, &user.Email, &user.PP, user.Birthday)
	fmt.Println("ta grosse darone", user)
	a, _ := json.Marshal(user)
	w.Write(a)

}
func cartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var cart Cart
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&cart)
	fmt.Println(cart)
	if cart.Pierre_ID != 0 {
		insert := `INSERT INTO labaiepierre.cart (pierre_ID, user_ID, pierre_price) VALUES ("` + strconv.Itoa(cart.Pierre_ID) + `","` + strconv.Itoa(cart.User_ID) + `","` + strconv.Itoa(cart.Pierre_price) + `")`
		fmt.Println(insert)
		db.Query(insert)
	}
	var query, _ = db.Query(`SELECT * FROM labaiepierre.cart WHERE USER_ID="` + strconv.Itoa(cart.User_ID) + `"`)
	var allCart []Cart
	for query.Next() {
		var cart Cart
		query.Scan(&cart.ID, &cart.User_ID, &cart.Pierre_ID, &cart.Pierre_price)
		allCart = append(allCart, cart)
	}
	fmt.Println(allCart)
	a, _ := json.Marshal(allCart)
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
