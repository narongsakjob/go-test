package main
import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My first golang")
	})
	http.HandleFunc("/user", user)
	http.HandleFunc("/product", product)
	http.ListenAndServe(":8080", nil)
}

func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Product Request")
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Request")
}