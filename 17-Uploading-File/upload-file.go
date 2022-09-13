package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

func renderFileUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		parseFile, _ := template.ParseFiles("templates/file-upload.html")
		parseFile.Execute(w, nil)
	} else {
		fmt.Fprintf(w, "Upload request")
	}
}

func handleFileUpload(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("error getting a file for the provided form key : ", err)
		return
	}
	defer file.Close()
	out, pathError := os.Create("/tmp/" + header.Filename)
	if pathError != nil {
		log.Printf("error creating a file for writing : ", pathError)
		return
	}
	defer out.Close()
	_, copyFileError := io.Copy(out, file)
	if copyFileError != nil {
		log.Printf("error occurred while file copy : ", copyFileError)
	}
	fmt.Fprintf(w, "File uploaded successfully : "+header.Filename)
}

func main() {
	http.HandleFunc("/", renderFileUpload)
	http.HandleFunc("/upload", handleFileUpload)

	fmt.Println("Starting server in port :", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil))
}
