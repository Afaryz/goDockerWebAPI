package main

import (
	"fmt"
	"net/http"
)

func main() {
	request1()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Printf("Welcome to the Go Web API")
	fmt.Println("Endpoint Hit: homepage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "Kevin Kowalski"

	fmt.Fprintf(response, "A little about...")
	fmt.Println("Endpoint Hit: ", who)

}
func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
}
