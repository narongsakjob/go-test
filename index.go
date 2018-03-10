package main
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

type Product struct {
	Name string
	Price int
}

func main() {
	var templates = template.Must(template.ParseFiles("index.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		myProduct := Product{"Kanom", 500}
		templates.ExecuteTemplate(w, "index.html", myProduct)
	})
	r.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "signup.html")
	})
	r.HandleFunc("/login", login)
	r.HandleFunc("/user/{name}", user).Methods("GET")
	r.HandleFunc("/product", product)
	http.ListenAndServe(":8080", r)

}

func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product Request")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method : ", r.Method)
	r.ParseForm()
	fmt.Println("Username : ", r.Form["username"])
	fmt.Println("Password : ", r.Form["password"])
	http.ServeFile(w, r, "login.html")
}

func user(w http.ResponseWriter, r *http.Request) {

	userDB:=map[string]int{
		"narongsak": 20,
		"test1": 25,
		"test2": 30,
	}

	vars := mux.Vars(r)
	name := vars["name"]
	age := userDB[name]
	fmt.Fprintf(w, "My name is %s %d years old", name, age)

}