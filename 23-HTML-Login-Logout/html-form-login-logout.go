package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

//creating a secure cookie, passing a hash key as the first argument, and a block key as the second argument.
//The hash key is used to authenticate values using HMAC and the block key is used to encrypt values.
var cookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

//get username from cookie
//userName defined as the return value
func getUserName(r *http.Request) (userName string) {
	//retrieve cookie with name session from request
	cookie, err := r.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		//decode cookie with name session into object cookieValue
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {
			//extract username from cookieValue object
			userName = cookieValue["username"]
		}
	}
	return userName
}

//set new cookie, with name as session
func setSession(username string, w http.ResponseWriter) {
	value := map[string]string{
		"username": username,
	}

	//encode cookie with username value
	encoded, err := cookieHandler.Encode("session", value)
	if err == nil {
		//generate cookie
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		//set generated cookie to ResponseWriter
		http.SetCookie(w, cookie)
	}
}

//clear cookie of encoded value
func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//Accept username and password
//Create new cookie if username and password are non empty, redirect to "/home" page
//if username or password are invalid redirect to "/"
func login(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")
	target := "/"
	if username != "" && password != "" {
		setSession(username, response)
		target = "/home"
	}
	http.Redirect(response, request, target, 302)
}

//clear cookie and redirect to landing page "/"
func logout(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

//Accessing the login page should redirect to home page if user is still logged on
func loginPage(w http.ResponseWriter, r *http.Request) {
	userName := getUserName(r)
	if userName != "" {
		http.Redirect(w, r, "/home", 302)
	} else {
		parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
		parsedTemplate.Execute(w, nil)
	}

}

func homePage(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		data := map[string]interface{}{
			"userName": userName,
		}
		parsedTemplate, _ := template.ParseFiles("templates/home.html")
		parsedTemplate.Execute(response, data)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/", loginPage)
	router.HandleFunc("/home", homePage)
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("POST")
	http.Handle("/", router)

	fmt.Println("Starting server in port:", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, router))
}
