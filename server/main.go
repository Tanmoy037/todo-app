package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/Tanmoy037/todo-app/router"
	   
)
func main() {
	r := router.Router()
	fmt.Println("Server running on port 3000...")

	log.Fatal(http.ListenAndServe(":3000", r))
}
