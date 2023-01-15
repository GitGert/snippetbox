package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {

	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	///////////
	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles("./ui/html/pages/home.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "internal server Error", http.StatusInternalServerError)
		return
	}
	// We then use the Execute() method on the template set to write the
	// template content as the response body. The last parameter to Execute()
	// represents any dynamic data that we want to pass in, which for now we'll
	// leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internalservererror", http.StatusInternalServerError)
	}
	///////////
	w.Write([]byte("Hello from Snppetbox"))
}

// add a snippetView ahnlder function
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to an integer using the strconv.Atoi() function. If it can't
	// be converted to an integer, or the value is less than 1, we return a 404 page
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 0 {
		// http.Error(w, "Can't find id", http.StatusNotFound)???
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id) // Frintf is okay to use to write strings like this
	// w.Write([]byte("display a spcific snippet...")) // this was here before
}

// add a snippetCreate hanlder function
func snippetCreate(w http.ResponseWriter, r *http.Request) {

	// Use r.Method to check whether the request is using POST or not.
	if r.Method != http.MethodPost { //THIS LINE IS THE SAME AS: r.Method != "POST"
		//////////////
		//use the Header().Set() method to add an 'Allow: POST' header to the
		//response header map. Te first parameter is the header name, and
		//the second parameter is the header value
		w.Header().Set("Allow", "POST")
		//////////////
		// If it's not, use the w.WriteHeader() method to send a 405 status
		// code and the w.Write() method to write a "Method Not Allowed"
		// response body. We then return from the function so that the
		// subsequent code is not executed.
		//////////////////////
		// instead of these two codes under this comment
		///// w.WriteHeader(405)
		///// w.Write([]byte("Method not allowed"))
		//DO THIS:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed) // the last one is the same as 405

		return
	}

	w.Write([]byte("Create a new snippet..."))
}
