package main
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My first golang")
	})
	http.HandleFunc("/user/", user)
	http.HandleFunc("/product", product)
	http.ListenAndServe(":8080", nil)

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