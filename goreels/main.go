package main

import (
	"log"
	"net/http"
	
)

func main() { // begin func main()

	http.Handle("/", http.FileServer(http.Dir("public"))) // pass in our public directory

	const addr =":8080"
	err := http.ListenAndServe(addr, nil)
	if (err != nil) {
		log.Fatalf("Server has failed: %v", err) // kills the server when an error is hit.
	} 

} // end func main() 
