package main

import (
	"fmt"
	"net/http"

)

 func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello, World!");
}
func aboutPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "This is the about page");
}


func main() {	
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloWorld)
	mux.HandleFunc("/about",aboutPage)

	fmt.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", mux)


	if err != nil {
		fmt.Println("Error starting server:", err)
	}else {		
		fmt.Println("Server started successfully")
	}




}