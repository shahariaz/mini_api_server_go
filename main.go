package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter,r *http.Request){
if r.URL.Path != "/hello"{
	http.Error(w,"404 not Found",http.StatusNotFound)
	return
}
if r.Method != "Get" {
	http.Error(w,"Method not supported",http.StatusNotFound)
	return
}
fmt.Fprintf(w,"Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request){
if err := r.ParseForm(); err != nil{
	fmt.Fprintf(w,"ParseForm failed: %v ",err)
	return
}
  fmt.Fprintf(w,"Form submitted with data: %v",r.Form)
	name := r.FormValue("name")
	address:= r.FormValue("address")
	fmt.Fprintf(w,"Name: %s\n",name)
	fmt.Fprintf(w,"Address: %s\n",address)
}
func main(){
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello",helloHandler)
	fmt.Printf("starting server at http://localhost:8080")
	if err:= http.ListenAndServe(":8000",nil); err != nil{
		log.Fatal(err)
	}
}