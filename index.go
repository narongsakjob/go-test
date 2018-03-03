package main
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My first golang")
	})
	r.HandleFunc("/user/", user)
	r.HandleFunc("/product", product)
	http.ListenAndServe(":8080", r)

}

func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product Request")
}

func user(w http.ResponseWriter, r *http.Request) {

	userDB:=map[string]int{
		"narongsak": 20,
		"test1": 25,
		"test2": 30,
	}

	name:=r.URL.Path[len("/user/"):]
	age:=userDB[name]
	fmt.Fprintf(w, "My name is %s %d years old", name, age)

}