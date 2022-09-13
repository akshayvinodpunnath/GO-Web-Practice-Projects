package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/schema"
)

const (
	CONNECT_HOST           = "localhost"
	CONNECT_PORT           = "8080"
	USERNAME_ERROR_MESSAGE = "Please enter valid username"
	PASSWORD_ERROR_MESSAGE = "Please enter valid password"
	GENERIC_ERROR_MESSAGE  = "Validation error"
)

type User struct {
	Username string `valid:"alpha,required"`
	Password string `valid:"alpha,required"`
}

func readForm(r *http.Request) *User {
	r.ParseForm()
	user := new(User)
	decoderErr := schema.NewDecoder().Decode(user, r.PostForm)
	if decoderErr != nil {
		log.Fatal("error mapping parsed form data", decoderErr.Error())
	}
	return user
}

func validateUser(w http.ResponseWriter, r *http.Request, user *User) (bool, string) {
	valid, validationError := govalidator.ValidateStruct(user)
	if !valid {
		usernameError := govalidator.ErrorByField(validationError, "Username")
		passwordError := govalidator.ErrorByField(validationError, "Password")
		if usernameError != "" {
			log.Printf("username validation error : ", usernameError)
			return valid, USERNAME_ERROR_MESSAGE
		}

		if passwordError != "" {
			log.Printf("password validation error : ", passwordError)
			return valid, PASSWORD_ERROR_MESSAGE
		}
	}
	return valid, GENERIC_ERROR_MESSAGE
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		parsefiles, _ := template.ParseFiles("templates/login-form.html")
		parsefiles.Execute(w, nil)
	} else {
		user := readForm(r)
		valid, validationErrorMessage := validateUser(w, r, user)
		if !valid {
			fmt.Fprintf(w, validationErrorMessage)
			return
		}
		fmt.Fprintf(w, "Hello "+user.Username+"!")
	}
}

func main() {
	http.HandleFunc("/", login)

	fmt.Println("Starting server on port :", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil))
}
